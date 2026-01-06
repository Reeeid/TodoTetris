"use client";

import { useState, useEffect } from "react";

interface Todo {
  id: number;
  subject: string; 
  description: string;
  uuid: string;
}

export default function TodoApp() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [loading, setLoading] = useState(true);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [editSubject, setEditSubject] = useState("");
  const [editDescription, setEditDescription] = useState("");

  const fetchTodos = async () => {
    try {
      const res = await fetch("/api/todo");
      if (res.ok) {
        const data = await res.json();
        setTodos(data.todos || []);
      }
    } catch (error) {
      console.error("Failed to fetch todos", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleEdit = (todo: Todo) => {
    setEditingId(todo.id);
    setEditSubject(todo.subject);
    setEditDescription(todo.description);
  };

  const cancelEdit = () => {
    setEditingId(null);
    setEditSubject("");
    setEditDescription("");
  };

  const saveEdit = async (todo: Todo) => {
    try {
      const res = await fetch("/api/todo", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ 
            uuid: todo.uuid,
            subject: editSubject,
            description: editDescription
        }),
      });

      if (!res.ok) throw new Error("Update failed");

      setTodos(todos.map(t => t.id === todo.id ? { ...t, subject: editSubject, description: editDescription } : t));
      setEditingId(null);
    } catch (e) {
      alert("Failed to update task");
      console.error(e);
    }
  };

  if (loading) return <div className="text-center p-10 text-gray-500 animate-pulse">Loading daily review...</div>;

  return (
    <div className="w-full max-w-4xl mx-auto p-6">
      <header className="mb-10 text-center">
        <h1 className="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-emerald-500 to-teal-600 mb-2">
          TODO
        </h1>
        <p className="text-gray-600 dark:text-gray-400">
          修正だけは可能です。
        </p>
      </header>

      <div className="grid gap-4 max-w-2xl mx-auto">
        {todos.length === 0 ? (
          <div className="text-center p-10 bg-gray-50 rounded-xl dark:bg-zinc-800 border-2 border-dashed border-gray-200 dark:border-zinc-700">
            <p className="text-xl text-gray-500">クリア済み</p>
          </div>
        ) : (
          todos.map((todo) => (
            <div
              key={todo.id}
              className="group p-5 bg-white rounded-xl shadow-sm border border-gray-100 dark:bg-zinc-800 dark:border-zinc-700 hover:shadow-md transition-all"
            >
              {editingId === todo.id ? (
                <div className="space-y-3">
                    <input 
                        className="w-full p-2 border rounded dark:bg-zinc-700 dark:border-zinc-600 dark:text-white"
                        value={editSubject}
                        onChange={e => setEditSubject(e.target.value)}
                    />
                    <textarea 
                        className="w-full p-2 border rounded h-20 dark:bg-zinc-700 dark:border-zinc-600 dark:text-white"
                        value={editDescription}
                        onChange={e => setEditDescription(e.target.value)}
                    />
                    <div className="flex justify-end gap-2">
                        <button onClick={cancelEdit} className="px-3 py-1 text-sm text-gray-500 hover:text-gray-700">キャンセル</button>
                        <button onClick={() => saveEdit(todo)} className="px-3 py-1 text-sm bg-emerald-500 text-white rounded hover:bg-emerald-600">保存</button>
                    </div>
                </div>
              ) : (
                <div className="flex items-start justify-between gap-4">
                    <div className="flex gap-4">
                        <span className="w-3 h-3 mt-2 rounded-full bg-emerald-500 flex-shrink-0" />
                        <div>
                            <h3 className="text-lg font-semibold text-gray-800 dark:text-gray-200">{todo.subject}</h3>
                            {todo.description && <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">{todo.description}</p>}
                        </div>
                    </div>
                    <button 
                        onClick={() => handleEdit(todo)}
                        className="text-sm text-indigo-500 hover:text-indigo-600 opacity-0 group-hover:opacity-100 transition-opacity"
                    >
                        修正
                    </button>
                </div>
              )}
            </div>
          ))
        )}
      </div>
    </div>
  );
}
