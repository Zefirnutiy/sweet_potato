
import { Link } from "react-router-dom"
import { InformationCard } from '../InformationCard/InformationCard';
import "./Navbar.css"
import "../../fonts/fontawesome-free-6.1.2-web/css/all.css"
import { FC } from "react";

type ILink = {
    icon: string
    title: string
    path: string
}

type PropTypes = {
    links: ILink[]
    organizationName?: string
    children?: React.ReactNode
  };

export const Navbar: FC<PropTypes> = ({links, organizationName='Wains', children}) =>{
    return (
        <div id="navbar">
        <div id="logo">
           {organizationName}
        </div>
        <div id="nav">
            {links?.map(link => 
                <Link to={link.path}><i className={link.icon}></i> {link.title}</Link> 
            )}
        </div>
        <div id="navContent">
            {children}
        </div>
    </div>
    )
}