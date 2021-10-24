
pragma solidity 0.4.17;


contract Decentify {

    struct Requests {

        string description;
        uint value;
        address recepient;
        bool isComplete;

    }

    address  public manager; 

    uint    public minContri;
    
    address[] public doners;

    function Decentify(uint argforMIN) public {
        manager = msg.sender;
        minContri = argforMIN;
    }

    function donate() public payable {

        require(msg.value > minContri);
        doners.push(msg.sender);
    }
 
}