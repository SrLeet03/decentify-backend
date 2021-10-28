import React, { useState , useEffect } from 'react'


import axois from 'axios' ; 

import { Row, Button } from 'react-bootstrap';

import FundCard from './FundCard'

//import './css/card.css'

const arr = [1, 1, 1, 1, 1, 1, 1];



export default function FundraiseList() {

    const [tag, setTag] = useState("");

    useEffect (() =>{
         
    } , [tag]) ; 
    const handleselect = (arg) => {
        setTag(arg);
    }

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
                            return <FundCard />;
                        })
                    }
                </Row>
            </div>


        </div>
    )
}
