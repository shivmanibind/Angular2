import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormControl ,ReactiveFormsModule} from '@angular/forms'
import { User } from '../model/user.component'
import { UserService } from '../service/service.component'
@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
users: User[] = [];
userForm= new FormGroup({
        email: new FormControl(),
        password: new FormControl(),
    });
    constructor(public data: UserService) { }

    ngOnInit() { } 

    onLoggedin() {
        localStorage.setItem('isLoggedin', 'true');
    }
onSubmit(){
    this.data.getAll().subscribe(users => { this.users = users; console.log("Users are " +this.users[0].name ); 
         });
    
}
}
