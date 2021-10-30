import React, { useState , useEffect } from 'react'


import axois from 'axios' ; 
import { useSelector  , useDispatch } from 'react-redux';
import { Row, Button } from 'react-bootstrap';

import FundCard from './FundCard'
import axios from 'axios';

//import './css/card.css'

const arr = [1, 1, 1, 1, 1, 1, 1];



export default function FundraiseList() {

    const [tag, setTag] = useState("");
    const dispatch = useDispatch() ;  
    useEffect (() =>{
        
        const data = axios.get(`${process.env.URL}/getallFundrs`)
        .then((result) =>{
            console.log(result) ; 
        }).catch((err) => {
            console.log(err) ; 
        })

    } , [tag]) ; 
    const handleselect = (arg) => {
        setTag(arg);
    }
     const fundrlist = useSelector((state) => state) ; 
     console.log("check" , fundrlist) ; 
    return (
        <div>
            all fundraisers
            <div className="tags">
                <b>select tag to donate your prefered category fundraiser : </b>
                <Button onClick = {() =>handleselect("all categories")} variant="info">All categories</Button>{'   '}
                <Button onClick ={() =>handleselect("medical")}variant="info">Medical</Button>{'   '}

                <Button onClick={() =>handleselect("enviroment")} variant="info">Enviroment</Button>{' '}

                <Button onClick={() =>handleselect("education")} variant="info">Education</Button>{' '}


                <Button onClick={() =>handleselect("social service")} variant="info">Social Service</Button>{' '}
                <br></br>
                {
                    tag!==""?<b>{tag} Fundraisers : </b> : ""
                }

            </div>
           
            <br></br>
            <br></br> <br></br>
            <div className="justify-content-md-center">
                <Row class="row ">
                    {
                        arr.map((key, value) => {
                            const prop = {
                                name  : 'sarvesh raut',
                                title : "Helps us please!",
                                date : '10 Nov,2019',
                                tagline : "Some quick example text to build on the card title and make up the bulk of the card's"
                            }
                            return <FundCard  info = {prop} />;
                        })
                    }
                </Row>
            </div>


        </div>
    )
}
