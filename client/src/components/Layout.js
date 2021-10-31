import React from 'react';

import { Switch, Route , useRouteMatch } from 'react-router-dom'

import FundraiseList from './blockchain/FundraiseList';
import LandingPage from './source/LandingPage';
import Notfound from './source/Notfound';

import FundProfile from './blockchain/FundProfile';
import Navbar from './Navbar';
import Privateroute from './auth/Privateroute';
import Profile from './auth/Profile';
import Signup from './auth/Signup';
const Layout = () => {

    let { url , path } = useRouteMatch();
        

    return (
        <>
        <Navbar/>
            <Switch >
                <Route  path='/'  exact component={LandingPage }  />
                <Route  path='/fundraisers' exact component={FundraiseList}/>
                <Route  path='/fundraisers/:id' exact component={FundProfile}/>
                <Route  path='/list' exact component={FundraiseList}/>
                <Route  path='/signup' exact component={Signup}/>
                <Privateroute  path='/myprofile'  component={Profile} isAuth = {true} />
                <Route component={Notfound}/>
            </Switch> 
           
        </>
    );

}

export default Layout;