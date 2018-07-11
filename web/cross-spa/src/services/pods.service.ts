import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from '../../node_modules/rxjs/internal/Observable';
import { environment } from '../environments/environment';

@Injectable()
export class PodService {

  constructor(private http: HttpClient) { }

  getPods(size_id?: number, flavor_id?: number): Observable<Pod[]> {
    return this.http.get(environment.apiRoot + '/product/pod', {
      params: {
        size_id: size_id ? size_id.toString() : '0',
        flavor_id: flavor_id ? flavor_id.toString() : '0'
      }
    }) as Observable<Pod[]>;
  }

  crossPods(pod_id: number): Observable<Pod[]> {
    return this.http.get(environment.apiRoot + '/cross/pod', {
      params: {
        pod_id: pod_id.toString()
      }
    }) as Observable<Pod[]>;
  }
}

export class Pod {
  pod_id: number;
  size_id: number;
  size_name: string;
  flavor_id: number;
  flavor_name: string;
  sku: string;
  quantity: number;
}
