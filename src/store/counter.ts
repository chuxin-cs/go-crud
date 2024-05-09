import {create} from 'zustand'

interface ICounterState {
    counter: number
    increase: (by: number) => void
}

const useCounterStore = create<ICounterState>((set) => ({
    counter: 0,
    increase: (by) => set((state) => ({counter: state.counter + by})),
}))


export default useCounterStore;