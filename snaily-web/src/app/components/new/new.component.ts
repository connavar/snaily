import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {ItemService} from '../../services/item.service';
import {ItemV1} from '../../model/item-v1';
import {UserState} from '../../model/state/userState';
import {HttpErrorResponse} from '@angular/common/http';
import {Messages} from '../../util/Messages';

declare var Swal: any;

@Component({
  selector: 'app-maintain',
  templateUrl: './new.component.html',
  styleUrls: ['./new.component.css']
})
export class NewComponent implements OnInit {

  private email: string;
  private body: string;
  private title: string;
  private day: number;
  private nmonth: number;
  private year: number;
  private hour: number;
  private minute: number;


  private theme = '';

  constructor(private router: Router, private itemService: ItemService) {
  }

  ngOnInit() {
    this.setTimeAsNow();
  }

  public request(): void {
    this.saveItem();
  }


  private saveItem(): void {

    const i = new ItemV1();
    i.title = this.title;
    i.body = this.body;
    i.createdBy = UserState.getMyID();
    i.waitingFor = this.email;


    this.itemService.post(i, (result) => (
        this.animateSuccess(result)
      ),
      (err: HttpErrorResponse) => {
        if (err.status === 403) {
          Messages.AccessDenied();
          this.router.navigate(['./login']);
          return;
        }
      }
    );
  }

  private animateSuccess(result: any) {
    console.log(result);
    Swal.fire({
      position: 'middle-end',
      type: 'success',
      title: 'Saved',
      showConfirmButton: false,
      timer: 1800
    });

    setTimeout(() => {
      window.location.href = './view?id=' + result.id;
      // this.router.navigate(['./view?id=' + result.id]); //TODO: use router
    }, 1500);

  }


  private setTimeAsNow() {

    const n = new Date();

    this.nmonth = n.getUTCMonth();
    this.year = n.getFullYear();

    this.day = n.getUTCDate();
    this.hour = n.getUTCHours();
    this.minute = n.getUTCMinutes();

  }
}
