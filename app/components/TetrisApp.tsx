"use client";

import { useState, useEffect, useCallback, useRef } from "react";

// --- Types ---
interface Todo {
  id: number;
  subject: string;
  uuid: string;
}

// Brought from TodoEntry
interface PendingTodo {
  id: string; 
  subject: string;
  description: string;
}

interface TetrisAppProps {
  onGameOver: () => void;
  pendingTodos?: PendingTodo[];
}

// Grid: 10 columns x 20 rows
const COLS = 10;
const ROWS = 20;

// Colors for pieces
const COLORS = [
  "bg-cyan-500",    // I
  "bg-blue-500",    // J
  "bg-orange-500",  // L
  "bg-yellow-500",  // O
  "bg-green-500",   // S
  "bg-purple-500",  // T
  "bg-red-500",     // Z
];

const SHAPES = [
  // I (4x4)
  [
    [0, 0, 0, 0],
    [1, 1, 1, 1],
    [0, 0, 0, 0],
    [0, 0, 0, 0]
  ],
  // J (3x3)
  [
    [1, 0, 0],
    [1, 1, 1],
    [0, 0, 0]
  ],
  // L (3x3)
  [
    [0, 0, 1],
    [1, 1, 1],
    [0, 0, 0]
  ],
  // O (2x2)
  [
    [1, 1],
    [1, 1]
  ],
  // S (3x3) - Note: standard SRS s is usually [0,1,1],[1,1,0].
  [
    [0, 1, 1],
    [1, 1, 0],
    [0, 0, 0]
  ],
  // T (3x3)
  [
    [0, 1, 0],
    [1, 1, 1],
    [0, 0, 0]
  ],
  // Z (3x3)
  [
    [1, 1, 0],
    [0, 1, 1],
    [0, 0, 0]
  ]
];

interface Cell {
  filled: boolean;
  color: string;
  uuid: string | null; // Todo UUID attached to this block
}

// --- Helper: Custom Hook for Game Loop ---
function useInterval(callback: () => void, delay: number | null) {
  const savedCallback = useRef(callback);
  useEffect(() => { savedCallback.current = callback; }, [callback]);
  useEffect(() => {
    if (delay !== null) {
      const id = setInterval(() => savedCallback.current(), delay);
      return () => clearInterval(id);
    }
  }, [delay]);
}

export default function TetrisApp({ onGameOver, pendingTodos }: TetrisAppProps) {
  // --- State ---
  const [grid, setGrid] = useState<Cell[][]>([]);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [activePiece, setActivePiece] = useState<{ x: number, y: number, shape: number[][], color: string, uuids: (string | null)[] } | null>(null);
  const [score, setScore] = useState(0);
  const [gameOver, setGameOver] = useState(false);
  const [dropSpeed, setDropSpeed] = useState<number | null>(1000);
  const [loadingMsg, setLoadingMsg] = useState<string | null>("読み込み中...");

  // Logic Hoisting
  const deserializeGrid = (jsonStr: string): Cell[][] => {
      const g = Array.from({ length: ROWS }, () => 
        Array.from({ length: COLS }, () => ({ filled: false, color: "", uuid: null }))
      );
      try {
          const cells = JSON.parse(jsonStr);
          if(Array.isArray(cells)) {
              cells.forEach((c: any) => {
                  if(c.y < ROWS && c.x < COLS) {
                      g[c.y][c.x] = { filled: true, color: c.color, uuid: c.uuid };
                  }
              });
          }
      } catch(e) { console.error("Bad board state", e); }
      return g;
  };

  const serializeGrid = (g: Cell[][]): string => {
      const filledCells = [];
      for(let y=0; y<ROWS; y++) {
          for(let x=0; x<COLS; x++) {
              if(g[y][x].filled) {
                  filledCells.push({
                      x, y, 
                      color: g[y][x].color,
                      uuid: g[y][x].uuid
                  });
              }
          }
      }
      return JSON.stringify(filledCells);
  };

  const checkCollision = (piece: any, currentGrid: Cell[][], moveX: number, moveY: number, newShape?: number[][]) => {
      const shape = newShape || piece.shape;
      if (!currentGrid || currentGrid.length === 0) return true;

      for (let y = 0; y < shape.length; y++) {
          for (let x = 0; x < shape[y].length; x++) {
              if (shape[y][x]) {
                  const newX = piece.x + x + moveX;
                  const newY = piece.y + y + moveY;
                  if (newX < 0 || newX >= COLS || newY >= ROWS) return true;
                  if (newY >= 0) {
                      if (!currentGrid[newY] || !currentGrid[newY][newX]) return true;
                      if (currentGrid[newY][newX].filled) return true;
                  }
              }
          }
      }
      return false;
  };

  // --- Initialization ---
  // --- Initialization ---
  const initializedRef = useRef(false);

  useEffect(() => {
    if (initializedRef.current) return;
    initializedRef.current = true;

    const init = async () => {
      let initialGrid: Cell[][] = Array.from({ length: ROWS }, () => 
        Array.from({ length: COLS }, () => ({ filled: false, color: "", uuid: null }))
      );
      
      try {
        const res = await fetch("/api/tetris");
        if (res.ok) {
          const data = await res.json();
          console.debug("[Tetris] Session Data:", data);

          // 1. Try to load existing board state
          if (data.session && data.session.board_state) {
              const savedGrid = deserializeGrid(data.session.board_state);
              // Simple check if valid grid (has correct dimensions)
              if (savedGrid && savedGrid.length === ROWS && savedGrid[0].length === COLS) {
                  console.debug("[Tetris] Loaded saved board state.");
                  setGrid(savedGrid);
                  setLoadingMsg(null);
                  return; // Skip penalty logic if we resumed a session
              }
          }

          // 2. If no saved state, apply New Game / Penalty Logic
          if (data.session && data.session.last_played_at) {
              const lastPlayed = new Date(data.session.last_played_at);
              const now = new Date();
              const diffTime = Math.abs(now.getTime() - lastPlayed.getTime());
              const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));
              
              if (diffDays > 0) {
                  const penaltyLines = Math.min(diffDays, 10);
                  console.debug(`[Tetris] Applying penalty: ${penaltyLines} lines (Days missed: ${diffDays})`);
                  
                  for (let i=0; i<penaltyLines; i++) {
                      for(let r=0; r < ROWS - 1; r++) {
                          initialGrid[r] = [...initialGrid[r+1]];
                      }
                      const garbageRow = Array.from({length: COLS}, () => ({
                          filled: Math.random() > 0.3,
                          color: "bg-gray-600",
                          uuid: null 
                      }));
                      garbageRow[Math.floor(Math.random() * COLS)].filled = false;
                      initialGrid[ROWS-1] = garbageRow;
                  }
              }
          }
        }
      } catch (e) { console.error(e); }

      setGrid(initialGrid);
      setLoadingMsg(null);
    };

    init();
  }, []);

  // --- Main Logic (Lock & Commit) ---
  const saveSession = async (currentGrid: Cell[][], currentScore: number) => {
      try {
          await fetch("/api/tetris", {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify({ 
                  Board_state: serializeGrid(currentGrid),
                  Score: currentScore
              }) 
          });
      } catch(e) { console.error("Save failed", e); }
  };

  const handleGameOver = async () => {
      setGameOver(true);
      setDropSpeed(null);
      // Game Over -> Delete ALL Todos (Including pending if created? No, only API known ones)
      // Since we create on lock, if game over happens BEFORE first lock, no backend change needed.
      // But if we fail middle of game?
      // Wait, "One Piece per Day" flow implies game ends after 1 piece.
      // So Game Over usually means "Failed to place piece" (Top out).
      // In that case, we should probably delete everything if requested "Game Over = Delete All".
      
      // If we haven't locked ANY piece yet, pending todos are just in memory. We just lose them.
      // If we are strictly 1 piece, then top out is unlikely unless penalty kills us.
      
      // Request: "Game Over -> Delete All"
      // If we have existing todos (from previous sessions? No, 1 piece flow).
      // Let's implement Delete All just in case.
      try {
         // Maybe fetch existing todos first to get UUIDs?
         const r = await fetch("/api/todo");
         const d = await r.json();
         if (d.todos && d.todos.length > 0) {
             const allIds = d.todos.map((t:any) => t.uuid);
             await fetch("/api/todo", {
                 method: "DELETE",
                 headers: { "Content-Type": "application/json" },
                 body: JSON.stringify({ uuids: allIds })
             });
         }
      } catch(e) {}

      // Save empty session
      await saveSession([], 0);
      alert("ゲームオーバー。タスクは全て削除されました。");
      onGameOver();
  };

  const lockingRef = useRef(false);

  const lockPiece = async () => {
      if (!activePiece || lockingRef.current) return;
      lockingRef.current = true;
      setDropSpeed(null);
      console.log("Locking piece...");

      // 1. Create Commit Payload
      // We will define the NEW grid first with Temp IDs
      let currentGrid = grid.map(row => row.map(cell => ({...cell})));
      let uuidIdx = 0;
      activePiece.shape.forEach((row, y) => {
          row.forEach((val, x) => {
               if(val) {
                   const u = activePiece.uuids[uuidIdx++]; 
                   const gy = activePiece.y + y;
                   const gx = activePiece.x + x;
                   if (gy >= 0 && gy < ROWS && gx >= 0 && gx < COLS) {
                       currentGrid[gy][gx] = {
                           filled: true,
                           color: activePiece.color,
                           uuid: u
                       };
                   }
               }
          });
      });

      // 2. Clear Lines (Calculated BEFORE UUID replacement to keep logic simple, 
      // but strictly we should check Extinguished *after* assigning real UUIDs?
      // Actually, if we clear lines, the UUIDs in those lines are GONE from the grid.
      // So we must capture them.
      
      // Strategy: 
      // A. Create Todos -> Get Real UUIDs
      // B. Map Temp->Real in the Grid.
      // C. Clear Lines.
      // D. Check what remains.
      
      // A & B: Batch Create & Map
      let createdUUIDs: string[] = [];
      if (pendingTodos && pendingTodos.length > 0) {
          try {
              const res = await fetch("/api/todo", {
                  method: "POST",
                  headers: { "Content-Type": "application/json" },
                  body: JSON.stringify({
                      todos: pendingTodos.map(p => ({
                          subject: p.subject,
                          description: p.description,
                          uuid: ""
                      }))
                  })
              });
              
              if (res.ok) {
                  const data = await res.json();
                  if (data.todos) {
                      const realList = data.todos.map((t: any) => t.uuid);
                      createdUUIDs = realList;
                      
                      // Map in Grid
                      currentGrid = currentGrid.map(row => row.map(cell => {
                          if (cell.uuid && cell.uuid.startsWith("TEMP_IDX_")) {
                               const idx = parseInt(cell.uuid.split("_")[2]);
                               if (realList[idx]) return { ...cell, uuid: realList[idx] };
                          }
                          return cell;
                      }));
                  }
              }
          } catch(e) {
              console.error("Creation failed", e);
              alert("サーバーエラーが発生しました。データが保存されていない可能性があります。");
              onGameOver();
              return;
          }
      }

      // C. Clear Lines && Scoring
      // First, capture ALL UUIDs currently on the board (Pre-Clear)
      const uuidsBeforeClear = new Set<string>();
      currentGrid.forEach(row => row.forEach(c => { if (c.uuid) uuidsBeforeClear.add(c.uuid); }));

      let linesCleared = 0;
      const clearedGrid = currentGrid.filter(row => {
          const isFull = row.every(cell => cell.filled);
          if (isFull) linesCleared++;
          return !isFull;
      });
      while (clearedGrid.length < ROWS) {
          clearedGrid.unshift(Array.from({ length: COLS }, () => ({ filled: false, color: "", uuid: null })));
      }
      
      const newScore = score + linesCleared;
      
      // D. Extinguished Check (Post-Clear)
      const uuidsAfterClear = new Set<string>();
      clearedGrid.forEach(row => row.forEach(c => { if (c.uuid) uuidsAfterClear.add(c.uuid); }));

      // Any UUID that was there BEFORE but is NOT there AFTER is extinguished
      const extinguished: string[] = [];
      uuidsBeforeClear.forEach(u => {
          if (!uuidsAfterClear.has(u)) {
              extinguished.push(u);
          }
      });

      if (extinguished.length > 0) {
          try {
              await fetch("/api/todo", {
                  method: "DELETE",
                  headers: { "Content-Type": "application/json" },
                  body: JSON.stringify({ uuids: extinguished })
              });
              console.log("Extinguished Tasks:", extinguished);
          } catch(e) { console.error(e); }
      }

      // Final: Save Session
      await saveSession(clearedGrid, newScore);
      
      alert(`セッション完了！ スコア: ${newScore}`);
      onGameOver();
  };


  // --- Game Loop components ---
  useInterval(() => {
      if (!activePiece || gameOver) return;
      if (!checkCollision(activePiece, grid, 0, 1)) {
          setActivePiece(prev => prev ? { ...prev, y: prev.y + 1 } : null);
      } else {
          lockPiece();
      }
  }, dropSpeed);

  // Spawn Effect
  const [hasSpawned, setHasSpawned] = useState(false);
  useEffect(() => {
      if (!activePiece && !gameOver && !loadingMsg && !hasSpawned) {
          // Spawn Logic
          // Pick Shape
          const typeIdx = Math.floor(Math.random() * SHAPES.length);
          const shape = SHAPES[typeIdx];
          const color = COLORS[typeIdx]; 
          
          let blockCount = 0;
          shape.forEach(row => row.forEach(cell => { if(cell) blockCount++; }));
          
          const pieceUUIDs: (string | null)[] = [];
          if (pendingTodos && pendingTodos.length > 0) {
              // Assign Temp IDs
               for(let i=0; i<blockCount; i++) {
                 const idx = i % pendingTodos.length; 
                 pieceUUIDs.push(`TEMP_IDX_${idx}`);
               }
          } else {
              // Empty play?
               for(let i=0; i<blockCount; i++) pieceUUIDs.push(null);
          }
          
           const startX = Math.floor((COLS - shape[0].length) / 2);
           const newPiece = { x: startX, y: 0, shape, color, uuids: pieceUUIDs };
           
           if (checkCollision(newPiece, grid, 0, 0)) {
               handleGameOver(); // Penalty killed us immediately
           } else {
               setActivePiece(newPiece);
               setHasSpawned(true); // Only spawn ONCE per session in strict flow?
                                    // User said "Play Tetris". Doesn't strictly say 1 piece.
                                    // But "Completed -> Request -> Delete -> Todo List ".
                                    // Usually implies 1 cycle.
                                    // If we want multiple pieces, we'd need to delay the commit until "Game Over" or "User Quit"?
                                    // Requirement: "Play is complete -> Request"
                                    // And "One piece per day" was previous context.
                                    // I'll stick to 1 piece spawn = Session. 
                                    // Wait, if I clear 1 line, I continue?
                                    // "Play completed -> Request".
                                    // If I spawn only once, then locking = completion.
                                    // So `hasSpawned` ensures only 1 piece.
           }
      }
  }, [activePiece, gameOver, loadingMsg, hasSpawned, pendingTodos, grid]);

  // Controls (Minimal)
  const move = (dir: number) => {
      if (!activePiece || gameOver) return;
      if (!checkCollision(activePiece, grid, dir, 0)) {
          setActivePiece(prev => prev ? { ...prev, x: prev.x + dir } : null);
      }
  };
  const rotate = (dir: 1 | -1) => {
       if (!activePiece || gameOver) return;
       const shape = activePiece.shape;
       const N = shape.length;
       
       // Matrix Rotate
       const newShape = Array.from({length:N}, () => Array(N).fill(0));
       if (dir === 1) { // CW
           for(let y=0; y<N; y++) {
               for(let x=0; x<N; x++) {
                   newShape[x][N-1-y] = shape[y][x];
               }
           }
       } else { // CCW
           for(let y=0; y<N; y++) {
               for(let x=0; x<N; x++) {
                   newShape[N-1-x][y] = shape[y][x];
               }
           }
       }

       // Wall Kick / Offset Test
       // Simple heuristic kicks: Center, Left, Right, Up, Diagonals
       const kicks = [
           {x:0, y:0}, 
           {x:-1, y:0}, {x:1, y:0},  // Shift Horizontal
           {x:0, y:-1},              // Shift Up (Floor kick)
           {x:-1, y:1}, {x:1, y:1},  // Down? (Unlikely but ok)
           {x:-2, y:0}, {x:2, y:0}   // I piece long moves
       ];

       for (let k of kicks) {
           if (!checkCollision(activePiece, grid, k.x, k.y, newShape)) {
               setActivePiece(prev => prev ? { 
                   ...prev, 
                   x: prev.x + k.x,
                   y: prev.y + k.y,
                   shape: newShape 
                   // rotationIndex update would go here if we tracked it
               } : null);
               return; // Success
           }
       }
  };
  const drop = () => {
      if (!activePiece || gameOver) return;
      let checkY = 0;
      while (!checkCollision(activePiece, grid, 0, checkY + 1)) checkY++;
      setActivePiece(prev => prev ? { ...prev, y: prev.y + checkY } : null);
  };


  if (loadingMsg) return <div className="text-white text-center p-10">{loadingMsg}</div>;

  return (
    <div className="flex flex-col items-center justify-center p-2 bg-zinc-900 min-h-screen text-white select-none touch-none">
       <div className="mb-4 text-xl font-bold tracking-widest text-transparent bg-clip-text bg-gradient-to-r from-yellow-400 to-orange-500">
           スコア: {score}
       </div>
       <div 
        className="relative bg-black border-4 border-zinc-700 overflow-hidden shadow-2xl mx-auto box-content"
        style={{ width: COLS * 25, height: ROWS * 25 }}
       >
         {grid.map((row, y) => row.map((cell, x) => (
             cell.filled ? (
                 <div 
                    key={`${y}-${x}`}
                    className={`absolute border border-black/20 ${cell.color}`}
                    style={{ left: x * 25, top: y * 25, width: 25, height: 25 }}
                 >
                    {cell.uuid && <div className="w-1 h-1 bg-white/50 rounded-full m-auto mt-2" />}
                 </div>
             ) : null
         )))}
         {activePiece && activePiece.shape.map((row, y) => row.map((val, x) => {
             if (val) {
                 const px = (activePiece.x + x) * 25;
                 const py = (activePiece.y + y) * 25;
                 return (
                  <div 
                     key={`p-${y}-${x}`}
                     className={`absolute border border-black/20 ${activePiece.color}`}
                     style={{ left: px, top: py, width: 25, height: 25 }}
                  />
                 );
             }
             return null;
         }))}
       </div>

       {/* Mobile-Friendly Control Cluster */}
       <div className="grid grid-cols-3 gap-4 mt-8 w-full max-w-xs px-4">
           {/* Top Row: Rotations centered or split? Let's put rotations on top corners */}
           <div className="flex flex-col items-center">
                <button onClick={() => rotate(-1)} className="w-16 h-16 bg-indigo-600 rounded-full shadow-lg active:scale-95 transition flex items-center justify-center">
                    <span className="text-2xl">↺</span>
                </button>
                <span className="text-xs text-gray-400 mt-1">左回転</span>
           </div>
           
           <div className="flex flex-col items-center pt-8"> {/* Offset Down button slightly */}
                <button onClick={drop} className="w-16 h-16 bg-gray-700 rounded-full shadow-lg active:scale-95 transition flex items-center justify-center">
                    <span className="text-2xl">↓</span>
                </button>
                <span className="text-xs text-gray-400 mt-1">落下</span>
           </div>

           <div className="flex flex-col items-center">
                <button onClick={() => rotate(1)} className="w-16 h-16 bg-indigo-600 rounded-full shadow-lg active:scale-95 transition flex items-center justify-center">
                    <span className="text-2xl">↻</span>
                </button>
                 <span className="text-xs text-gray-400 mt-1">右回転</span>
           </div>

           {/* Bottom Row: Movement */}
           <div className="flex flex-col items-center">
               <button onClick={() => move(-1)} className="w-16 h-16 bg-gray-700 rounded-full shadow-lg active:scale-95 transition flex items-center justify-center">
                   <span className="text-2xl">←</span>
               </button>
               <span className="text-xs text-gray-400 mt-1">移動</span>
           </div>

            {/* Empty center or maybe a hold button later? For now empty filler */}
           <div className="flex items-center justify-center pointer-events-none opacity-0">
               <div className="w-16 h-16" />
           </div>

           <div className="flex flex-col items-center">
               <button onClick={() => move(1)} className="w-16 h-16 bg-gray-700 rounded-full shadow-lg active:scale-95 transition flex items-center justify-center">
                   <span className="text-2xl">→</span>
               </button>
               <span className="text-xs text-gray-400 mt-1">移動</span>
           </div>
       </div>
    </div>
  );
}
