import { combineReducers } from "redux";

import { fundrsReducer , fundrReducer } from "./reducer"; 

export const reducers =  combineReducers({
     allFundrs : fundrsReducer,
     fundr : fundrReducer
});