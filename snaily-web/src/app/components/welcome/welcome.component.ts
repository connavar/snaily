import {Component, OnInit} from '@angular/core';
import {UserState} from '../../model/state/userState';
import {Router} from '@angular/router';

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {

  constructor(private router: Router) {
  }

  ngOnInit() {

    if (UserState.isLoggedIn()) {
      const link = ['./home'];
      this.router.navigate(link);
    }

  }

}
