
pragma solidity 0.4.17;


contract Decentify {

    struct Request {

        string description;
        uint value;
        address recepient;
        bool isComplete;
        uint apprCount;
        mapping(address => bool)approvers;
    }

    address  public manager; 
    uint    public minContri;
    mapping( address=>bool) public doners;
    uint  public totLen;
    Request[] public requests;
    
    modifier accessRestricted() {
        
        require(msg.sender == manager);
        _;
    }

    modifier accessApprove() {
        
        require(doners[msg.sender] == true);
        _;
    }

    function Decentify(uint argforMIN) public {
        manager = msg.sender;
        minContri = argforMIN;
        totLen = 0;
    }

    function donate() public payable {

        require(msg.value > minContri);
        doners[msg.sender] = true;
        totLen++;
    }

    function createRequest(string des, uint value, address recp) 
    public accessRestricted {
        Request memory newReq  =  Request(des, value, recp, false, 0);
        //here we don't need to inilize appovers map caz this is refr type dataType 
        //and all others are value type.
        requests.push(newReq);
    }

    function approveRequest(uint index) public accessApprove {

        Request storage curRequest = requests[index]; 

        require(curRequest.approvers[msg.sender] == false);
        curRequest.apprCount++;
        curRequest.approvers[msg.sender] = true;
    }

    function greenflagForReq(uint index) public accessRestricted {
        Request storage curRequest = requests[index];
        require(curRequest.isComplete == false);
        require((totLen/2) < curRequest.apprCount);
        curRequest.isComplete = true;
        curRequest.recepient.transfer(curRequest.value);
        curRequest.isComplete = true;
        
    }
 

}