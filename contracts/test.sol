
pragma solidity 0.8.9;


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