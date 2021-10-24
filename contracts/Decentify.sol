
pragma solidity 0.4.17;


contract Decentify {

    struct Request {

        string description;
        uint value;
        address recepient;
        bool isComplete;

    }

    address  public manager; 
    uint    public minContri;
    mapping( address=>bool) public doners;
    Request[] public requests;
    
    modifier accessRestricted() {
        
        require(msg.sender == manager);
        _;
    }

    function Decentify(uint argforMIN) public {
        manager = msg.sender;
        minContri = argforMIN;
    }

    function donate() public payable {

        require(msg.value > minContri);
        doners[msg.sender] = true;
    }

    function createRequest(string des, uint value, address recp) 
    public accessRestricted {
        Request memory newReq  =  Request(des, value, recp, false);
        requests.push(newReq);
    }
 

}