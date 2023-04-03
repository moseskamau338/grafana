import { AppChromeService } from './AppChromeService';

describe('AppChromeService', () => {
  it('onToggleKioskMode should set chromeless to true when searchbar is hidden', () => {
    const chromeService = new AppChromeService();
    chromeService.onToggleSearchBar();
    chromeService.onToggleKioskMode();
    expect(chromeService.state.getValue().chromeless).toBe(true);
  });

  it('Ignore state updates when sectionNav and pageNav have new instance but same text, url or active child', () => {
    const chromeService = new AppChromeService();
    let stateChanges = 0;

    chromeService.state.subscribe(() => stateChanges++);
    chromeService.update({ sectionNav: { text: 'hello' }, pageNav: { text: 'test', url: 'A' } });
    chromeService.update({ sectionNav: { text: 'hello' }, pageNav: { text: 'test', url: 'A' } });

    expect(stateChanges).toBe(2);

    // if url change we should update
    chromeService.update({ sectionNav: { text: 'hello' }, pageNav: { text: 'test', url: 'new/url' } });
    expect(stateChanges).toBe(3);

    // if active child changed should update state
    chromeService.update({
      sectionNav: { text: 'hello' },
      pageNav: { text: 'test', url: 'A', children: [{ text: 'child', active: true }] },
    });
    expect(stateChanges).toBe(4);

    // If active child is the same we should not update state
    chromeService.update({
      sectionNav: { text: 'hello' },
      pageNav: { text: 'test', url: 'A', children: [{ text: 'child', active: true }] },
    });
    expect(stateChanges).toBe(4);
  });

  it('still updates if pageNav parent text changes so breadcrumbs can update', () => {
    const chromeService = new AppChromeService();
    let stateChanges = 0;

    chromeService.state.subscribe(() => stateChanges++);
    expect(stateChanges).toBe(1);

    chromeService.update({ sectionNav: { text: 'hello' }, pageNav: { text: 'test', url: 'A' } });
    expect(stateChanges).toBe(2);

    // if parent change we should update
    chromeService.update({
      sectionNav: { text: 'hello' },
      pageNav: { text: 'test', url: 'new/url', parentItem: { text: 'parent' } },
    });
    expect(stateChanges).toBe(3);

    // if a higher parent changes we should still update
    chromeService.update({
      sectionNav: { text: 'hello' },
      pageNav: { text: 'test', url: 'new/url', parentItem: { text: 'parent', parentItem: { text: 'grandparent' } } },
    });
    expect(stateChanges).toBe(4);
  });
});
