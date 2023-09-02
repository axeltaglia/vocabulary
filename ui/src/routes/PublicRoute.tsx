import {Navigate} from "react-router-dom"

export type PublicRouteProps = {
    isAuthenticated: boolean
    outlet: JSX.Element
};

export default function PublicRoute({isAuthenticated, outlet}: PublicRouteProps) {
    if(!isAuthenticated) {
        return outlet
    } else {
        return <Navigate to={{ pathname: '/home' }} />
    }
};