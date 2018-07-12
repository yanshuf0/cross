import { Component } from '@angular/core';
import { Pod, PodService } from '../services/pods.service';
import { MachineService, CoffeeMachine } from '../services/machines.service';

interface SizeCategory {
  size_id: number;
  size_name: string;
  cols: number;
  rows: number;
}
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  podSrc = [
    'assets/coffee-pod1.jpeg',
    'assets/coffee-pod2.jpeg',
    'assets/coffee-pod3.jpeg',
    'assets/coffee-pod4.jpeg'
  ];
  itemSelected: boolean;
  machineSizes: SizeCategory[] = [];
  podSizes: SizeCategory[] = [];
  pods: Pod[] = [];
  machines: CoffeeMachine[] = [];
  showMachines: boolean;
  showPods: boolean;
  assetFolder = '/assets';
  crossPods: Pod[] = [];
  selectedPod: Pod;
  selectedMachine: CoffeeMachine;
  selectingMachines: boolean;
  selectingPods: boolean;
  displayMachines: CoffeeMachine[];
  displayMachineSizes: boolean;
  displayPods: Pod[];
  showDisplayPods: boolean;
  constructor(
    private podService: PodService,
    private machineService: MachineService
  ) {}
  async getPods() {
    this.clearSet();
    this.selectingPods = true;
    this.pods = await this.podService.getPods().toPromise();
    for (const p of this.pods) {
      if (!this.podSizes.find(s => s.size_id === p.size_id)) {
        this.podSizes.push({
          size_id: p.size_id,
          size_name: p.size_name,
          cols: 1,
          rows: 1
        });
      }
    }
    this.showPods = true;
  }
  async getMachines() {
    this.clearSet();
    this.machines = await this.machineService.getMachines().toPromise();
    for (const m of this.machines) {
      if (!this.machineSizes.find(s => s.size_id === m.size_id)) {
        this.machineSizes.push({
          size_id: m.size_id,
          size_name: m.size_name,
          cols: 1,
          rows: 1
        });
      }
    }
    this.showMachines = true;
  }
  async getCrossPods(pod_id?: number, machine_id?: number) {
    if (pod_id) {
      this.crossPods = await this.podService.crossPods(pod_id).toPromise();
    }
    if (machine_id && !pod_id) {
      this.crossPods = await this.machineService
        .crossMachines(machine_id)
        .toPromise();
    }
  }

  async selectPod(pod_id: number) {
    this.pods = await this.podService.getPods().toPromise();
    this.selectedPod = this.pods.find(p => p.pod_id === pod_id);
    this.selectedMachine = null;
    this.itemSelected = true;
    this.getCrossPods(pod_id);
  }

  async selectMachine(machine_id: number) {
    this.machines = await this.machineService.getMachines().toPromise();
    this.selectedMachine = this.machines.find(
      m => m.coffee_machine_id === machine_id
    );
    this.itemSelected = true;
    this.getCrossPods(null, machine_id);
  }

  async selectMachineSize(size_id: number) {
    this.displayMachineSizes = true;
    this.selectingMachines = false;
    this.displayMachines = await this.machineService
      .getMachines(size_id)
      .toPromise();
  }

  async selectPodSize(size_id: number) {
    this.displayPods = await this.podService.getPods(size_id).toPromise();
    this.showDisplayPods = true;
    this.showPods = false;
  }

  clearSet() {
    this.displayMachines = null;
    this.displayPods = null;
    this.displayMachineSizes = false;
    this.crossPods = null;
    this.itemSelected = false;
    this.machines = null;
    this.machineSizes = [];
    this.pods = null;
    this.podSizes = [];
    this.selectedMachine = null;
    this.selectedPod = null;
    this.selectingMachines = false;
    this.selectingPods = false;
    this.showDisplayPods = false;
    this.showDisplayPods = false;
    this.showPods = false;
    this.showMachines = false;
  }
}
