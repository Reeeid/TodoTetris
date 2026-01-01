"use client";

import { useEffect, useState } from "react";

export default function Home() {
  const [uuid, setUuid] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch("http://localhost:3000/api/uuid",{method: "GET"})
      .then((res) => {
        if (!res.ok) {
          throw new Error(`Status: ${res.status}`);
        }
        return res.json();
      })
      .then((data) => {
        console.log("API Data:", data);
        setUuid(data.uuid);
      })
      .catch((err) => {
        console.error("Fetch error:", err);
        setError(err.message);
      });
  }, []);

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-gray-100 p-4">
      <main className="flex w-full max-w-lg flex-col items-center gap-8 rounded-xl bg-white p-8 shadow-lg">
        <h1 className="text-2xl font-bold">API Test</h1>
        
        <div className="text-center">
          <p className="mb-2 text-gray-600">Requesting: <code className="bg-gray-200 px-1 rounded">/api/UUID</code></p>
          
          {error ? (
            <div className="rounded-md bg-red-50 p-4 text-red-600">
              <p className="font-bold">Error:</p>
              <p>{error}</p>
            </div>
          ) : uuid ? (
            <div className="rounded-md bg-green-50 p-4 text-green-700">
              <p className="font-bold">Success! Received UUID:</p>
              <p className="break-all font-mono text-xl">{uuid}</p>
            </div>
          ) : (
            <p className="animate-pulse text-gray-500">Loading...</p>
          )}
        </div>
      </main>
    </div>
  );
}
