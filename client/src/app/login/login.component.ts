import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from '../app.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginInstance:LoginModel = new LoginModel();

  constructor(private service:AppService, private router: Router) { }

  ngOnInit() {
  }

  loginToSystem() {
    this.service.loginToSystem(this.loginInstance).subscribe(data => {
      AppService.TOKEN = data.token;
      this.router.navigateByUrl('/home');
    })
  }
}

export class LoginModel {
  username:string;
  password:string;
}
