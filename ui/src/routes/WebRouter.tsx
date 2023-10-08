import {
    BrowserRouter as Router,
    Routes,
    Route,
} from 'react-router-dom';
import {ReactElement} from "react";
import { routes as appRoutes } from "./routes";
import PublicRoute from "./PublicRoute";

function Path() {
    const defaultPrivateRouteProps = {
        isAuthenticated: false,
        authenticationPath: '/',
    };

    return (
        <Router>
            <Routes>
                {appRoutes.map((route) => {
                    return <Route key={route.key} path={route.path}
                                  element={<PublicRoute {...defaultPrivateRouteProps} outlet={<route.component/>}/>}/>
                })}
            </Routes>
        </Router>
    ) as ReactElement
}

export default Path
