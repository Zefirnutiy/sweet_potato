import {Routes, Route, Navigate} from 'react-router-dom'
import {LoginPage} from "./pages/LoginPage"
import {RegisterPage} from "./pages/RegisterPage"


export const useRoutes: React.FC<boolean> = (isAuthenticted: boolean) => {

    if(isAuthenticted){
        return (
            <h1>Ok</h1>
        )
    }else{
       return ( 
       <Routes>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="*" element={<Navigate to="/login" />} />
        </Routes>
        )
    }


}