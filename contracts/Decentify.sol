
pragma solidity 0.8.9;


contract FunraiseFactory {
      
   
    struct Fundra {

        string id;
        address addoffunr;
    }

    Funraiser[] public fundraisers; 
    mapping(Funraiser => string) public funrId;

    function createFunraiser(uint minAmo, string memory id) public {
        Funraiser curradd = new Funraiser(minAmo, msg.sender);
        funrId[curradd] = id; 
        fundraisers.push(curradd);
    }

    function getallFundraisers() view public returns (Funraiser[] memory) {

    return fundraisers;
}
    
}


contract Funraiser {

    

    struct Request {

        string description;
        uint value;
        address recepient;
        bool isComplete;
        uint apprCount;
        mapping(address => bool) approvers;
    }

    address  public manager; 
    uint    public minContri;
    mapping( address=>bool) public doners;
    uint  public totLen;
    Request[] public requests;
    
    constructor (uint val , address creator) public {
        manager = creator;
        minContri = val;
    }
    modifier accessRestricted() {
        
        require(msg.sender == manager);
        _;
    }

    modifier accessApprove() {
        
        require(doners[msg.sender] == true);
        _;
    }

    function donate() public payable {
        require(msg.value > minContri);
        doners[msg.sender] = true;
        totLen++;
    }

    function createRequest(string memory des, uint value, address recp) 
    public accessRestricted {
       // Request storage newReq  =  Request(des, value, recp, false, 0);
        Request storage newRequest = requests.push();
        newRequest.description = des;
        newRequest.value = value;
        newRequest.isComplete = false;
        newRequest.apprCount = 0;
        newRequest.recepient = recp;
        //here we don't need to inilize appovers map caz this is refr type dataType 
        //and all others are value type.
        //requests.push(newReq);
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
        payable(curRequest.recepient).transfer(curRequest.value);
        curRequest.isComplete = true;

    }
 

}