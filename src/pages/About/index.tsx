import {useState} from "react"
import {Button} from "antd";

const About = () => {
    console.log("我进入了about页面")
    const [total, setTotal] = useState(0)
    let num = 0
    const setNum = () => {
        num = num + 1;
        console.log(num)
    }

    return <>
        <h1>about</h1>
        <Button onClick={() => setTotal(total + 1)}>{total}</Button>
        <Button onClick={() => setNum()}>{num}</Button>
    </>
}
export default About
