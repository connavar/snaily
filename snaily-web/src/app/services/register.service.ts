import {Injectable} from '@angular/core';
import {MessageRegisterRequestV1} from '../model/register-request-v1';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  constructor(private http: HttpClient) {
  }

  public post(item: MessageRegisterRequestV1, win: Function, fail: Function): void {
    this.http.post(window.location.origin + environment.urlRegisterV1, item)
      .toPromise()
      .then((result) => {
          win(result);
        }
      )
      .catch((error) => {
        fail(error);
        console.log(error);
      });

  }

}
