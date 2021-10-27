import React from 'react';

import { Switch, Route , useRouteMatch } from 'react-router-dom'

import FundraiseList from './blockchain/FundraiseList';
import LandingPage from './source/LandingPage';
import Notfound from './source/Notfound';

import Navbar from './Navbar';

const Layout = () => {

    let { url , path } = useRouteMatch();
        

    return (
        <>
        <Navbar/>
            <Switch >
                <Route  path='/'  exact component={LandingPage }  />
                <Route  path='/fundraisers' exact component={FundraiseList}/>
                <Route  path='/list' exact component={FundraiseList}/>
                <Route component={Notfound}/>
            </Switch> 
           
        </>
    );

}

export default Layout;