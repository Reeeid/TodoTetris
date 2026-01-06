"use client";

import { useState } from "react";

export interface PendingTodo {
  id: string;
  subject: string;
  description: string;
}

interface TodoEntryProps {
  onStartGame: (tasks: PendingTodo[]) => void;
}

const PlusIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
);
const TrashIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
);
const PlayIcon = () => (
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><polygon points="5 3 19 12 5 21 5 3"/></svg>
);

export default function TodoEntry({ onStartGame }: TodoEntryProps) {
  const [pending, setPending] = useState<PendingTodo[]>([]);
  const [subject, setSubject] = useState("");
  const [description, setDescription] = useState("");

  const addPending = (e: React.FormEvent) => {
      e.preventDefault();
      if(!subject.trim()) return;
      const newItem = { 
          id: Date.now().toString() + Math.random().toString(), 
          subject, 
          description 
      };
      setPending([...pending, newItem]);
      setSubject("");
      setDescription("");
  };
  
  const removePending = (id: string) => {
      setPending(pending.filter(p => p.id !== id));
  };

  const handleStart = () => {
      if(pending.length === 0) return;
      onStartGame(pending);
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-4 md:p-8">
      <div className="text-center mb-10">
        <h2 className="text-4xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500 animate-gradient-x mb-4">
          タスク入力
        </h2>
        <p className="text-gray-500 dark:text-gray-400 text-lg max-w-lg mx-auto">
          今日やるべきことを登録してください。これらがテトリスのブロックになります。
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        
        <div className="bg-white dark:bg-zinc-800/50 backdrop-blur-xl border border-white/20 dark:border-zinc-700/50 p-6 rounded-3xl shadow-xl">
          <form onSubmit={addPending} className="flex flex-col gap-5">
            <div>
              <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">タスク名</label>
              <input
                type="text"
                value={subject}
                onChange={(e) => setSubject(e.target.value)}
                placeholder="例: メール返信、買い物..."
                className="w-full p-4 bg-gray-50 dark:bg-zinc-900/50 border border-gray-200 dark:border-zinc-700 rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none transition-all dark:text-white"
              />
            </div>
            <div>
               <label className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1">詳細 (任意)</label>
               <textarea
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="詳細やメモがあれば入力..."
                className="w-full p-4 h-32 bg-gray-50 dark:bg-zinc-900/50 border border-gray-200 dark:border-zinc-700 rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none transition-all dark:text-white resize-none"
              />
            </div>
            
            <button
              type="submit"
              disabled={!subject.trim()}
              className="mt-2 py-4 px-6 bg-zinc-900 dark:bg-white text-white dark:text-zinc-900 font-bold rounded-xl hover:scale-[1.02] active:scale-[0.98] transition-transform flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <PlusIcon />
              リストに追加
            </button>
          </form>
        </div>

        <div className="bg-white dark:bg-zinc-800/50 backdrop-blur-xl border border-white/20 dark:border-zinc-700/50 p-6 rounded-3xl shadow-xl flex flex-col h-full">
            <h3 className="text-xl font-bold text-gray-800 dark:text-white mb-4 flex items-center gap-2">
                <span>登録済みタスク</span>
                <span className="bg-purple-100 dark:bg-purple-900 text-purple-600 dark:text-purple-300 text-xs px-2 py-1 rounded-full">{pending.length}</span>
            </h3>
            
            <div className="flex-1 overflow-y-auto space-y-3 dark:text-gray-200 min-h-[200px] max-h-[400px] custom-scrollbar">
                {pending.length === 0 && (
                     <div className="h-full flex flex-col items-center justify-center text-gray-400 border-2 border-dashed border-gray-200 dark:border-zinc-700 rounded-2xl p-8">
                         <p>タスクはまだありません。</p>
                     </div>
                )}
                {pending.map((p) => (
                    <div key={p.id} className="group p-4 bg-gray-50 dark:bg-zinc-900/80 rounded-2xl flex items-start gap-3 border border-transparent hover:border-purple-500/30 transition-all">
                        <div className="flex-1">
                            <h4 className="font-semibold text-gray-900 dark:text-gray-100">{p.subject}</h4>
                            {p.description && <p className="text-sm text-gray-500 dark:text-gray-400 mt-1 line-clamp-2">{p.description}</p>}
                        </div>
                        <button
                          onClick={() => removePending(p.id)}
                          className="p-2 text-gray-400 hover:text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
                        >
                            <TrashIcon />
                        </button>
                    </div>
                ))}
            </div>

            <div className="mt-6 pt-6 border-t border-gray-100 dark:border-zinc-700">
               <button
                  onClick={handleStart}
                  disabled={pending.length === 0}
                  className="w-full py-5 bg-gradient-to-r from-purple-600 to-indigo-600 text-white font-bold text-lg rounded-2xl shadow-lg shadow-purple-500/30 hover:shadow-purple-500/50 hover:-translate-y-1 active:scale-[0.99] transition-all flex items-center justify-center gap-3 disabled:filter disabled:grayscale disabled:cursor-not-allowed"
                >
                    <span>テトリスを開始</span>
                    <PlayIcon />
                </button>
            </div>
        </div>
      </div>
    </div>
  );
}
