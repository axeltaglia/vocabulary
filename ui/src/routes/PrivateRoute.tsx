import {Navigate} from "react-router-dom"

export type PrivateRouteProps = {
    isAuthenticated: boolean
    authenticationPath: string
    outlet: JSX.Element
}

export default function PrivateRoute({isAuthenticated, authenticationPath, outlet}: PrivateRouteProps) {
    if(isAuthenticated) {
        return outlet
    } else {
        return <Navigate to={{ pathname: authenticationPath }} />
    }
};