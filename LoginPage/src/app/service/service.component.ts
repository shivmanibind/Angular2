import { Injectable } from '@angular/core';
import { Http, Headers, RequestOptions, Response } from '@angular/http';

//import { User } from '../_models/index';

@Injectable()
export class UserService {
     
    constructor(private http: Http) { 
       
    }

    getAll() {
       return this.http.get('http://localhost:12345/data').map((response: Response) => response.json());
    }
    
 postData(regForm) {
           var i=1
console.log("Inside Service",regForm.value)
    let headers = new Headers({ 'Content-Type': 'application/x-www-form-urlencoded' });
    let options = new RequestOptions({ headers: headers });
    let body = JSON.stringify(regForm.value);
   console.log("HIIIIIIII", body)
    return this.http.post('http://localhost:12345/data/8', body, options ).map((res: Response) => res.json());


     /*  return this.http.get('http://localhost:12345/data/i').map((response: Response) => response.json()); */
       
    }

        }