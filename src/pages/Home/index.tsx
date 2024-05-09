import useCounterStore from '@/store/counter'
import {Button} from "antd"

export default function Home() {
    const counter = useCounterStore((state) => state.counter)
    const increase = useCounterStore((state) => state.increase)

    return <>
        <h1>home</h1>
        <Button onClick={() => increase(1)}> counter: {counter} </Button>
    </>
}