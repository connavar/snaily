import {Injectable} from '@angular/core';
import {ItemV1} from '../model/item-v1';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {UserState} from '../model/state/userState';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  constructor(private http: HttpClient) {

  }

  public post(item: ItemV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlNewItemV1, item, {
        headers: new HttpHeaders().append('key', UserState.getMyKey())
      }
    )
      .toPromise()
      .then((result) => {
        win(result);
      })
      .catch((err: HttpErrorResponse) => {
        if (err.status === 403) {
          UserState.logout();
        }
        fail(err);
      });

  }

  public getMyItems(win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlMyItemsV1, {}, {
        headers: new HttpHeaders().append('key', UserState.getMyKey())
      }
    ).toPromise()
      .then((result) => {
        win(result);
      })
      .catch((err: HttpErrorResponse) => {
        if (err.status === 403) {
          UserState.logout();
        }
        fail(err);
      });
  }

  public getItem(item: ItemV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlGetItemV1, item, {
        headers: new HttpHeaders().append('key', UserState.getMyKey())
      }
    ).toPromise()
      .then((result) => {
        win(result);
      })
      .catch((err: HttpErrorResponse) => {
        if (err.status === 403) {
          UserState.logout();
        }
        fail(err);
      });
  }


}
