import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl ,ReactiveFormsModule} from '@angular/forms'
import { FormsModule } from '@angular/forms';
import { UserService } from '../service/service.component'
import { User } from '../model/user.component'

@Component({
    selector: 'app-signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.scss'],
})
export class SignupComponent implements OnInit {
 users: User[] = [];
    regForm= new FormGroup({
        name: new FormControl(),
        email: new FormControl(),
        password: new FormControl(),
        resetpwd: new FormControl(),
    });

    constructor( public data: UserService) { 


    }


    ngOnInit() { 
      
       
    }

   onSubmit(){
            this.data.postData(this.regForm).subscribe(data => {
                alert('ok');
          }, error => {
              console.log("Error");
          });;
        
    }


}
