import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from '../../node_modules/rxjs/internal/Observable';
import { environment } from '../environments/environment';
import { Pod } from './pods.service';

@Injectable()
export class MachineService {

  constructor(private http: HttpClient) { }

  getMachines(size_id?: number): Observable<CoffeeMachine[]> {
    return this.http.get(environment.apiRoot + '/product/machine', {
      params: {
        size_id: size_id ? size_id.toString() : '0'
      }
    }) as Observable<CoffeeMachine[]>;
  }

  crossMachines(cofee_machine_id: number): Observable<Pod[]> {
    return this.http.get(environment.apiRoot + '/cross/machine', {
      params: {
        coffee_machine_id: cofee_machine_id.toString()
      }
    }) as Observable<Pod[]>;
  }
}

export class CoffeeMachine {
  coffee_machine_id: number;
  size_id: number;
  size_name: string;
  sku: string;
  model_id: string;
  model_name: string;
  water_line: string;
}
