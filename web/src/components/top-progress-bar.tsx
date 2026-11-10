"use client";

import { useEffect, useState } from "react";
import { usePathname, useSearchParams } from "next/navigation";

export function TopProgressBar() {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const [loading, setLoading] = useState(false);
  const [progress, setProgress] = useState(0);

  // When route changes, finish the animation
  useEffect(() => {
    if (loading) {
      setProgress(100);
      const timer = setTimeout(() => {
        setLoading(false);
        setProgress(0);
      }, 300);
      return () => clearTimeout(timer);
    }
  }, [pathname, searchParams]);

  // Intercept internal link clicks to start loading animation
  useEffect(() => {
    const handleClick = (e: MouseEvent) => {
      const target = (e.target as HTMLElement).closest("a");
      if (!target) return;

      const href = target.getAttribute("href");
      if (!href) return;

      // Ignore external links, anchor hashes, new tabs, and modifier keys
      if (
        href.startsWith("http://") ||
        href.startsWith("https://") ||
        href.startsWith("#") ||
        href.startsWith("mailto:") ||
        target.target === "_blank" ||
        e.ctrlKey ||
        e.metaKey ||
        e.shiftKey ||
        e.altKey
      ) {
        return;
      }

      // Check if target href is different from current path
      const currentUrl = window.location.pathname + window.location.search;
      if (href !== currentUrl) {
        setLoading(true);
        setProgress(25);
        setTimeout(() => setProgress(65), 100);
        setTimeout(() => setProgress(85), 400);
      }
    };

    document.addEventListener("click", handleClick, { capture: true });
    return () => document.removeEventListener("click", handleClick, { capture: true });
  }, []);

  if (!loading && progress === 0) return null;

  return (
    <div
      style={{
        position: "fixed",
        top: 0,
        left: 0,
        height: "3px",
        width: `${progress}%`,
        zIndex: 999999,
        background: "linear-gradient(90deg, #3b82f6, #8b5cf6, #ec4899)",
        boxShadow: "0 0 12px rgba(59, 130, 246, 0.8), 0 0 4px rgba(236, 72, 153, 0.8)",
        transition: progress === 100 ? "width 0.2s ease-out, opacity 0.3s ease-out" : "width 0.4s cubic-bezier(0.4, 0, 0.2, 1)",
        opacity: progress === 100 ? 0 : 1,
        pointerEvents: "none",
      }}
    />
  );
}
