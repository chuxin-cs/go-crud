import {Suspense} from "react"

export default function lazyLoad(Component: React.LazyExoticComponent<() => JSX.Element>) {
    return (
        <Suspense>
            <Component/>
        </Suspense>
    )
}