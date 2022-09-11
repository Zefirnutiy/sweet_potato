import {createContext} from 'react'


interface IAuthContext{
    token: string | null
    userId: number | null
    login: (jwtToken: string, id: number) => null,
    logout: () => void
    isAuthenticated: boolean
}

export const AuthContext = createContext({} as IAuthContext)