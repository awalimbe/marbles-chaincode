package main

import (
  "fmt"
)


type loc struct  {
        requester_name string;
        beneficiary_name  string;
        amount  string;
        expiry_date string;
       	status string;
        advising_bank string;
        document_hash string;
        loc_filename string;
        contract_hash string;
        bol_hash string;
    }
   
  var counter uint = 0;
  var LOCs map[uint]*loc;
  
// Adding LOCs 
func addLoc(requester_name, beneficiary_name, amount, expiry_date, document_hash,loc_filename, contract_hash,  bol_hash string) (bool){
  
     counter = counter+1;
     LOCs[counter] =  &loc{};
     LOCs[counter].requester_name = requester_name;
     LOCs[counter].beneficiary_name= beneficiary_name;
     LOCs[counter].amount= amount;
     LOCs[counter].expiry_date= expiry_date;
     LOCs[counter].status= "requested";
     LOCs[counter].advising_bank = "none";
     LOCs[counter].document_hash= document_hash;
     LOCs[counter].loc_filename= loc_filename;
     LOCs[counter].contract_hash= contract_hash;
     LOCs[counter].bol_hash = bol_hash ;
     //LOCs [counter]= loc{requester_name,beneficiary_name,amount,expiry_date,"requested","none",document_hash,loc_filename,contract_hash, bol_hash};
     fmt.Println(LOCs [counter].requester_name);
     return true;

}

// Return all LOCs in the system
    func getLoc(location uint) []byte{
	b := make([]byte, 300)
             tracker:= 0;

        
      	for  i:=0 ; i<len(LOCs [location].requester_name) ; i++{
 		
		//fmt.Println(LOCs[location].requester_name[i]);
		b[tracker] = LOCs[location].requester_name[i] ;
	        tracker = tracker + 1;
         
           }

		  tracker = tracker + 1;
 	for  j:=0 ; j<len(LOCs [location].beneficiary_name) ; j++{
 		
		//fmt.Println(LOCs[location].beneficiary_name[j]);
		b[tracker] = LOCs[location].beneficiary_name[j] ;
	        tracker = tracker + 1;
         
           }
      
		  tracker = tracker + 1;

	 for  k:=0 ; k<len(LOCs [location].amount) ; k++{
 		
		//fmt.Println(LOCs[location].amount[k]);
		b[tracker] = LOCs[location].amount[k] ;
	        tracker = tracker + 1;
         
           }
		
		  tracker = tracker + 1;
 	for  l:=0 ; l<len(LOCs [location].expiry_date) ; l++{
 		
		//fmt.Println(LOCs[location].expiry_date[l]);
		b[tracker] = LOCs[location].expiry_date[l] ;
	        tracker = tracker + 1;
         
           }

		  tracker = tracker + 1;
 	for m:= 0; m <len(LOCs [location].status) ; m++{
            b[tracker] = LOCs[location].status[m] ;
            tracker = tracker + 1;
        }

	  tracker = tracker + 1;
 	for n:= 0; n <len(LOCs [location].advising_bank) ; n++{
            b[tracker] = LOCs[location].advising_bank[n] ;
            tracker = tracker + 1;
        }

	  tracker = tracker + 1;

	 for p:= 0; p <len(LOCs [location].document_hash); p++{
	     	//fmt.Println(LOCs[location].document_hash[p]);
            b[tracker] = LOCs[location].document_hash[p] ;
            tracker = tracker + 1;
        }
	
 	 tracker = tracker + 1;
		
	 for q:= 0; q <len(LOCs [location].loc_filename); q++{
	     	//fmt.Println(LOCs[location].loc_filename[q]);
            b[tracker] = LOCs[location].loc_filename[q] ;
            tracker = tracker + 1;
        }
		
	  tracker = tracker + 1;
		
	for r:= 0; r <len(LOCs [location].contract_hash); r++{
	     	//fmt.Println(LOCs[location].contract_hash[r]);
            b[tracker] = LOCs[location].contract_hash[r] ;
            tracker = tracker + 1;
        }

  	tracker = tracker + 1;
	
	for s:= 0; s <len(LOCs [location].bol_hash); s++{
	     	//fmt.Println(LOCs[location].bol_hash[s]);
            b[tracker] = LOCs[location].bol_hash[s] ;
            tracker = tracker + 1;
        }

		
        
                  return b;
        
    }


 //Get number of LOCs in the system
    func getNumberOfLocs () uint{
        return counter;
    }
    

//Update the advising bank & create LOC
    func createLoc (location uint,advising_bank string ) bool{
        LOCs[location].advising_bank = advising_bank;
        LOCs[location].status = "Created";
        return true;
    }

//Change LOC status to advised
    func adviseLoc (location uint) bool{
        LOCs[location].status = "Advised";
	fmt.Println(LOCs[location].status);
        return true;
    }

 //Change LOC status to indicate BOL uploaded
    func uploadBol (location uint, bol_hash string, status string) bool{
        LOCs[location].bol_hash = bol_hash;
        LOCs[location].status = status;
        return true;
    }

 //Change LOC status to indicate new Contract uploaded
    func uploadContract (location uint, contract_hash string, status string) bool{
        LOCs[location].contract_hash = contract_hash;
        LOCs[location].status = status;
        return true;
    }

  //Change LOC status to other statuses
    func updateLocStatus (location uint, status string) bool{
        LOCs[location].status = status;
        return true;
    }


func main(){

 LOCs = make(map[uint]*loc) //creating a map


 s:= addLoc("Shruthi","Harika","200", "12-12-2017", "dsfsdf", "dfsdfsdf", "xcbcbxc","sdfsdfs");
 w:= addLoc("Vinay","Rekha","100", "12-12-2017", "sdfsdfsd", "sjfhjdkfh", "sgsdgdsg","sdfsdfs");
 fmt.Println("Add LOC", s,w);

 z:=getLoc(1);
 fmt.Println("Get All LOC", z);

 b:= getNumberOfLocs();
 fmt.Println("Number of LOCs:",b);

 a:=createLoc (1,"aBank");
 fmt.Println("Create LOC:", a);
 
 d:=adviseLoc (1);
 fmt.Println("AdviseLOC:", d);

 e:=uploadBol(2,"BolHash","BOL_Uploaded");
 fmt.Println("BOL_Uploaded Status:", e);

 f:=uploadContract(1,"Contract","Contract_Uploaded");
 fmt.Println("Contract_UploadedStatus:", f);

 g:=updateLocStatus (2,"DealCompleted");
 fmt.Println("Update_Status:", g);

}
