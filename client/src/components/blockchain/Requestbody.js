import React from 'react'

import './css/profile.css'

export default function FundProfile({ info }) {
    return (
        <div>

            
                    <h1>Current Reuests : </h1>
                    <div className="job-body">
                        
                       
                        <ul>
                            <li>Courage - The courage to do the right thing every day</li>
                            <li>Ownership - We personally own each commitment and issue we touch</li>
                            <li>Respect - Respect for each team member will result in respect for all our stakeholders</li>
                            <li>Excellence - Average is never good enough - what difference did you make today.</li>
                        </ul>
                        <h2>Important Notes</h2>
                        <p>We look to positively impact the lives we touch by making a difference each day. Change drives our business each and every day and our culture allows us to manage and embrace change by establishing core values:</p>
                        <div className="details">
                            <p className="job-payrange"><strong>Pay</strong>: $13 to $17.50/hour</p>
                            <p className="date"><strong>Posted</strong>: <span itemprop="datePosted">04/29/2014</span></p>
                            <p className="job-status"><strong>Job Status</strong>: <span itemprop="employmentType">Full Time, Part Time</span></p>
                        </div>

                    </div>

                </div>
            
        
    )
}
