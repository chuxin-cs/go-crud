import { create } from 'zustand'
import { persist } from 'zustand/middleware'

interface IGlobalState {
    primaryColor: string
    setColor: (color: string) => void
}

const useGlobalStore = create<IGlobalState>()(
    persist(
        (set) => ({
            primaryColor: '#00b96b',
            setColor: (color) => set(() => ({ primaryColor: color })),
        }),
        {
            name: 'primaryColor',
            partialize: (state) =>
                Object.fromEntries(Object.entries(state).filter(([key]) => ['primaryColor'].includes(key))),
        }
    )
)

export default useGlobalStore

