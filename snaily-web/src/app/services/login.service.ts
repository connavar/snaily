import {Injectable} from '@angular/core';
import {MessageLoginRequestV1} from '../model/login-request-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) {
  }

  public post(item: MessageLoginRequestV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlLoginV1, item)
      .toPromise()
      .then((result) => {
        win(result);
      })
      .catch((error) => {
        fail(error);
      });

  }


}
