import { useState, useMemo } from "react";

function useBoolean(value: any) {
  const [state, setState] = useState(value);
  const actions = useMemo(() => {
    console.log(111);
    return {
      toggle: () => {
        setState(!state);
      },
      set: (v: any) => {
        setState(v);
      },
      setTrue: () => {
        setState(true);
      },
      setFalse: () => {
        setState(false);
      },
    };
  }, []);

  return [state, actions];
}

function App() {
  const [count, setCount] = useState(0);
  const [state] = useBoolean(true);
  const [state1] = useBoolean(true);

  return (
    <>
      <h1>
        Vite + React {state} = {state1}
      </h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
      </div>
    </>
  );
}

export default App;
