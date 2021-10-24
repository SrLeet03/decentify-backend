
pragma solidity 0.4.17;


contract FunraiseFactory {
      
   
    struct Fundra {

        string id;
        address addoffunr;
    }

    address[] public fundraisers; 
    mapping(address => string) public funrId;

    function createFunraiser(uint minAmo, string id) public {
        address curradd = new Funraiser(minAmo, msg.sender);
        funrId[curradd] = id; 
        fundraisers.push(curradd);
    }

    // function getallFundraisers() view public returns (address[]) {

    //    return fundraisers;
    // }
    
}


contract Funraiser {

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

    function Funraiser(uint argforMIN, address owner) public {
        manager = owner;
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