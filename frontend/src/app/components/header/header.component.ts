import { Component, HostListener } from '@angular/core';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {
  isMobileView: boolean = false;
  isDropdownOpen: boolean = false;

  @HostListener('window:resize', ['$event'])
  onResize(event: any) {
    this.checkWindowSize();
  }

  ngOnInit() {
    this.checkWindowSize();
  }

  toggleDropdown(): void {
    this.isDropdownOpen = !this.isDropdownOpen;
  }

  private checkWindowSize(): void {
    this.isMobileView = window.innerWidth <= 750;
  }
}
