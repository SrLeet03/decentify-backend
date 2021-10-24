
pragma solidity ^0.4.17;


contract Test {

    string public message; 

    function setMes(string memory newmsg)  public
    {
        message = newmsg;
    } 

    function getMsg()  public view returns ( string memory) {
        return message; 
    }

}