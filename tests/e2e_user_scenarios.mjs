/**
 * MakerSpell Forum - Automated End-to-End User Scenario Test Suite
 * 
 * Verifies live user scenarios, API contracts, 3-level printer hierarchy,
 * central MakerSpell ID authentication redirect, and security boundaries.
 * 
 * Usage:
 *   node forum/tests/e2e_user_scenarios.mjs [--url https://forum.makerspell.com]
 */

import http from 'http';
import https from 'https';
import { URL } from 'url';

const args = process.argv.slice(2);
const urlArgIdx = args.indexOf('--url');
const BASE_URL = (urlArgIdx !== -1 && args[urlArgIdx + 1]) 
  ? args[urlArgIdx + 1].replace(/\/$/, '') 
  : (process.env.FORUM_URL || 'https://forum.makerspell.com');

console.log(`\n======================================================`);
console.log(`  MakerSpell Forum - Automated E2E Scenario Tests`);
console.log(`  Target: ${BASE_URL}`);
console.log(`  Timestamp: ${new Date().toISOString()}`);
console.log(`======================================================\n`);

function request(endpoint, options = {}) {
  return new Promise((resolve, reject) => {
    const fullUrl = new URL(endpoint.startsWith('http') ? endpoint : `${BASE_URL}${endpoint}`);
    const client = fullUrl.protocol === 'https:' ? https : http;
    const startTime = Date.now();

    const req = client.request(fullUrl, {
      method: options.method || 'GET',
      headers: {
        'User-Agent': 'MakerSpell-E2E-Tester/1.0',
        ...(options.headers || {})
      }
    }, (res) => {
      let data = '';
      res.on('data', chunk => data += chunk);
      res.on('end', () => {
        const latencyMs = Date.now() - startTime;
        if (data.charCodeAt(0) === 0xFEFF) data = data.slice(1); // strip BOM
        let json = null;
        try {
          json = JSON.parse(data);
        } catch (_) {}
        resolve({
          statusCode: res.statusCode,
          headers: res.headers,
          body: data,
          json: json,
          latencyMs
        });
      });
    });

    req.on('error', (err) => reject(err));
    if (options.body) {
      req.write(typeof options.body === 'string' ? options.body : JSON.stringify(options.body));
    }
    req.end();
  });
}

let passed = 0;
let failed = 0;

function assert(condition, scenarioName, message) {
  if (condition) {
    console.log(`  ✅ [PASS] ${scenarioName}: ${message}`);
    passed++;
  } else {
    console.error(`  ❌ [FAIL] ${scenarioName}: ${message}`);
    failed++;
  }
}

async function runScenario1_SiteAndAuth() {
  console.log(`Scenario 1: Site Metadata & System Health`);
  const res = await request('/api/get-site');
  assert(res.statusCode === 200, 'Scenario 1', `HTTP status 200 (got ${res.statusCode})`);
  assert(res.json && res.json.status === 'ok', 'Scenario 1', `API status ok`);
  assert(res.latencyMs < 2000, 'Scenario 1', `Response latency acceptable (${res.latencyMs}ms)`);

  const authProviders = res.json?.data?.authProviders || [];
  const makerSpellAuth = authProviders.find(p => p.name === 'makerspell');
  assert(!!makerSpellAuth, 'Scenario 1', `MakerSpell ID auth provider registered`);
  assert(makerSpellAuth?.displayName === 'MakerSpell ID', 'Scenario 1', `Provider display name is 'MakerSpell ID'`);
}

async function runScenario2_CategoryHierarchy() {
  console.log(`\nScenario 2: Category Hierarchy & Navigation`);
  const res = await request('/api/get-categories');
  assert(res.statusCode === 200, 'Scenario 2', `HTTP status 200 (got ${res.statusCode})`);
  assert(res.json && res.json.status === 'ok', 'Scenario 2', `API status ok`);

  const categories = res.json?.data || [];
  assert(categories.length >= 7, 'Scenario 2', `Categories loaded (count: ${categories.length})`);

  const hardware = categories.find(c => c.slug === 'hardware');
  assert(!!hardware, 'Scenario 2', `Hardware category found (slug: hardware)`);

  // Check key brands
  const brandSlugs = ['creality', 'snapmaker', 'bambu-lab', 'voron', 'elegoo', 'anycubic', 'prusa'];
  for (const bSlug of brandSlugs) {
    const brand = categories.find(c => c.slug === bSlug);
    assert(!!brand, 'Scenario 2', `Brand category found: ${bSlug}`);
    if (brand && hardware) {
      assert(brand.parentId === hardware.id, 'Scenario 2', `Brand ${bSlug} is child of Hardware`);
    }
  }
}

async function runScenario3_AuthRedirection() {
  console.log(`\nScenario 3: MakerSpell Central Auth Redirection (SSO)`);
  const res = await request('/api/auth/start?provider=makerspell');
  assert(res.statusCode === 302, 'Scenario 3', `HTTP status 302 Redirect (got ${res.statusCode})`);

  const location = res.headers['location'] || '';
  assert(location.startsWith('https://id.makerspell.com/api/oauth/authorize'), 'Scenario 3', `Redirects to id.makerspell.com OAuth`);
  assert(location.includes('client_id=casnode'), 'Scenario 3', `Contains client_id=casnode`);
  assert(location.includes('code_challenge='), 'Scenario 3', `Includes PKCE code_challenge`);
  assert(location.includes('code_challenge_method=S256'), 'Scenario 3', `Uses S256 PKCE method`);
}

async function runScenario4_GuestSession() {
  console.log(`\nScenario 4: Guest Session State & Privacy`);
  const res = await request('/api/get-account');
  assert(res.statusCode === 200, 'Scenario 4', `HTTP status 200 (got ${res.statusCode})`);
  assert(res.json?.data === null, 'Scenario 4', `Guest account session is null without credentials`);
}

async function runScenario5_TopicDiscussions() {
  console.log(`\nScenario 5: Topic Discussions & Reading Feed`);
  const res = await request('/api/get-topics');
  assert(res.statusCode === 200, 'Scenario 5', `HTTP status 200 (got ${res.statusCode})`);
  assert(res.json && res.json.status === 'ok', 'Scenario 5', `API status ok`);

  const topics = res.json?.data?.topics || [];
  assert(topics.length > 0, 'Scenario 5', `Topic feed contains active topics (count: ${topics.length})`);
  if (topics.length > 0) {
    const firstTopic = topics[0];
    assert(typeof firstTopic.id === 'number', 'Scenario 5', `Topic has valid numeric ID: ${firstTopic.id}`);
    assert(typeof firstTopic.title === 'string' && firstTopic.title.length > 0, 'Scenario 5', `Topic has title: "${firstTopic.title}"`);
  }
}

async function runScenario6_SecurityBoundaries() {
  console.log(`\nScenario 6: Security Boundaries & Unauthenticated Write Protection`);
  const postRes = await request('/api/add-topic', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: { title: 'Unauthorized Test Topic', raw: 'Should be rejected', categoryId: 1 }
  });
  // Must reject unauthenticated write (status error or non-200)
  const isRejected = postRes.statusCode !== 200 || postRes.json?.status === 'error';
  assert(isRejected, 'Scenario 6', `Unauthenticated topic creation safely blocked (status: ${postRes.statusCode}, msg: "${postRes.json?.msg || ''}")`);

  const commentRes = await request('/api/add-post', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: { topicId: 1, raw: 'Unauthorized comment' }
  });
  const isCommentRejected = commentRes.statusCode !== 200 || commentRes.json?.status === 'error';
  assert(isCommentRejected, 'Scenario 6', `Unauthenticated post creation safely blocked (status: ${commentRes.statusCode}, msg: "${commentRes.json?.msg || ''}")`);
}

async function runScenario7_LatencyBenchmark() {
  console.log(`\nScenario 7: Performance & Latency Benchmark`);
  const samples = [];
  for (let i = 0; i < 3; i++) {
    const res = await request('/api/get-site');
    samples.push(res.latencyMs);
  }
  const avg = Math.round(samples.reduce((a, b) => a + b, 0) / samples.length);
  assert(avg < 1200, 'Scenario 7', `Average response latency is healthy: ${avg}ms (samples: ${samples.join(', ')}ms)`);
}

async function main() {
  try {
    await runScenario1_SiteAndAuth();
    await runScenario2_CategoryHierarchy();
    await runScenario3_AuthRedirection();
    await runScenario4_GuestSession();
    await runScenario5_TopicDiscussions();
    await runScenario6_SecurityBoundaries();
    await runScenario7_LatencyBenchmark();

    console.log(`\n------------------------------------------------------`);
    console.log(`  E2E Test Results: ${passed} passed, ${failed} failed (Total: ${passed + failed})`);
    console.log(`------------------------------------------------------\n`);

    if (failed > 0) {
      process.exit(1);
    }
  } catch (err) {
    console.error(`\n❌ Fatal test runner error:`, err);
    process.exit(1);
  }
}

main();
