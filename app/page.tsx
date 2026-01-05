"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import TodoEntry, { PendingTodo } from "./components/TodoEntry";
import TetrisApp from "./components/TetrisApp";
import TodoApp from "./components/TodoApp";

type ViewMode = "loading" | "entry" | "game" | "list";

export default function Dashboard() {
  const [mode, setMode] = useState<ViewMode>("loading");
  const [sessionTodos, setSessionTodos] = useState<PendingTodo[]>([]);
  const router = useRouter();

  useEffect(() => {
    checkStatus();
  }, []);

  const checkStatus = async () => {
    try {
      const res = await fetch("/api/tetris");
      
      if (res.status === 401) {
        router.push("/login");
        return;
      }

      if (!res.ok) {
        // Fallback or error
        console.error("Status check failed");
        return;
      }

      const data = await res.json();
      
      // Check if played today
      if (data.is_played) {
          setMode("list");
      } else {
          // If not played, go to entry (even if new user)
          setMode("entry");
      }

    } catch (error) {
      console.error(error);
    }
  };

  const handleStartGame = (todos: PendingTodo[]) => {
    setSessionTodos(todos);
    setMode("game");
  };

  const handleGameOver = () => {
    // After game, ALWAYS go to list
    setMode("list");
  };

  if (mode === "loading") {
    return (
      <div className="flex h-screen items-center justify-center bg-gray-50 dark:bg-zinc-900 flex-col gap-4">
        <div className="text-xl font-semibold text-gray-500 animate-pulse">
          Connecting to System...
        </div>
        {/* If it takes too long, we might want to show a retry manually? 
            For now, let's just make the text friendlier. 
            User complaint: "Shows up occasionally". 
            If it's just slow, this text is fine. 
        */}
      </div>
    );
  }

  // Add Error Mode Handling if needed, but for now let's make checkStatus robust.

  return (
    <div className="min-h-screen bg-gray-100 dark:bg-zinc-900 text-gray-800 dark:text-gray-100 font-sans">
      <nav className="p-4 bg-white dark:bg-zinc-800 shadow-sm flex justify-between items-center">
        <h1 className="text-xl font-bold tracking-tight">TodoTetris</h1>
        <button
          onClick={() => {
            document.cookie = "token=; Max-Age=0; path=/;";
            router.push("/login");
          }}
          className="text-sm font-medium text-red-500 hover:text-red-600"
        >
          Logout
        </button>
      </nav>

      <main className="container mx-auto py-10 px-4">
        {mode === "entry" && <TodoEntry onStartGame={handleStartGame} />}
        {mode === "game" && <TetrisApp onGameOver={handleGameOver} pendingTodos={sessionTodos} />}
        {mode === "list" && <TodoApp />}
      </main>
    </div>
  );
}
