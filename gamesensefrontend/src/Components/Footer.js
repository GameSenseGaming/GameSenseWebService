import React from 'react';
import "../style/componentsStyle/Footer.css";

const Footer = () => {
    return(
        
    <div className="main-footer">
        <div className='container'>
            <div className='row'>
                 <div>
                        <ul>
                                <li><a href="https://discord.com/channels/271429360081305601/876618747148910623">Discord</a></li>
                                <li><a href="https://github.com/ThomasCardin/Gamesense">GithubRepo/WebApp</a></li>
                                <li><a href="https://github.com/ThomasCardin/GameSenseBot">GithubRepo/GamesenseBot</a></li>
                                <li><a href="https://github.com/ThomasCardin/GameSenseBot">Code of conduct</a></li>
                                <li><a href="https://github.com/ThomasCardin/GameSenseBot">Download GSG Bot</a></li>
                         </ul>
                </div>
            </div>      
        </div>      
    </div>
    )   
}

export default Footer;