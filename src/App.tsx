import {RouterProvider} from 'react-router-dom'
import router from './router'
import {ConfigProvider} from 'antd'
import zhCN from 'antd/locale/zh_CN'
import {useGlobalStore} from '@/store'


import dayjs from "dayjs"
import 'dayjs/locale/zh-cn'

dayjs.locale('zh-cn')

function App() {
    const {primaryColor} = useGlobalStore()

    return (
        <ConfigProvider
            locale={zhCN}
            theme={{
                token: {
                    colorPrimary: primaryColor,
                },
            }}
        >
            <RouterProvider router={router}/>
        </ConfigProvider>
    );
}

export default App;
