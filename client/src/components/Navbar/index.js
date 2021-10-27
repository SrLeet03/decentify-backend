import React from 'react';
import {
Nav,
NavLink,
Bars,
NavMenu,
NavBtn,
NavBtnLink,
} from './NavbarElements';


import { FaAffiliatetheme } from "react-icons/fa";
import { FaAmazonPay } from "react-icons/fa";
import { FaBlogger } from "react-icons/fa";

import { FaFirefoxBrowser } from "react-icons/fa";

    

const Navbar = () => {
return (
	<>
	<Nav>
		<Bars />

		<NavMenu>
		<NavLink exact to='/' activeStyle>
			<FaAffiliatetheme/>Home
		</NavLink>
		<NavLink to='/about' activeStyle>
			<FaFirefoxBrowser/>Events
		</NavLink>
		<NavLink to='/fundraisers' activeStyle>
			<FaAmazonPay/>'   'Fundraisers
		</NavLink>
		
		<NavLink to='/blogs' activeStyle>
			<FaBlogger/>Blogs
		</NavLink>
		</NavMenu>
		<NavBtn>
		<NavBtnLink to='/signup'>Sign In</NavBtnLink>
		</NavBtn>
	</Nav>
	</>
);
};

export default Navbar;
