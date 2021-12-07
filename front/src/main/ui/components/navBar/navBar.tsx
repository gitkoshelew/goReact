import React from 'react';
import s from './navBar.module.css';
import {AppBar, Box, Button, IconButton, Toolbar, Typography} from '@mui/material';
import logo from '../../../../assets/img/AppBar/logo.png';
import login from '../../../../assets/img/AppBar/login.png';
import ourCompany from '../../../../assets/img/AppBar/ourCompany.png';
import home from '../../../../assets/img/AppBar/home.png';
import hotel from '../../../../assets/img/AppBar/hotel.png';
import {PATH} from '../../Routes/RoutesInfo';
import {NavLink} from 'react-router-dom';

const {logoTitle, btnGroup, btnOnly} = s;

export const NavBar = () => {
    return (

        <Box sx={{flexGrow: 1}}>
            <AppBar color={'transparent'} position="static">
                <Toolbar>
                    <IconButton
                        size="large"
                        edge="start"
                        color="inherit"
                        aria-label="menu"
                        sx={{mr: 2}}
                    >
                    </IconButton>
                    <Typography variant="h6" component="div" sx={{flexGrow: 1}}>
                        <div className={logoTitle}>
                            <img src={logo} alt={'logo'}/>
                        </div>
                    </Typography>
                    <div className={btnGroup}>
                        <NavLink to={PATH.HOME}><Button color="inherit">
                            <div className={btnOnly}>
                                <img src={home} alt="homePage"/>
                                <span>Home</span>
                            </div>
                        </Button>
                        </NavLink>
                        <NavLink to={PATH.ABOUT_US}>
                            <Button color="inherit">
                                <div className={btnOnly}>
                                    <img src={ourCompany} alt="ourCompany"/>
                                    <span>About us</span>
                                </div>
                            </Button>
                        </NavLink>
                        <NavLink to={PATH.HOTELS}><Button color="inherit">
                            <div className={btnOnly}>
                                <img src={hotel} alt="hotel"/>
                                <span>Hotel</span>
                            </div>
                        </Button>
                        </NavLink>
                        <NavLink to={PATH.LOGIN}><Button color="inherit">
                            <div className={btnOnly}>
                                <img src={login} alt="login"/>
                                <span>Login</span>
                            </div>
                        </Button>
                        </NavLink>
                    </div>
                </Toolbar>
            </AppBar>
        </Box>
    );
}
