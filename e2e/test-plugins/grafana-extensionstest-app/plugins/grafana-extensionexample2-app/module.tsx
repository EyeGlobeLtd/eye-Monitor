import { AppPlugin } from '@grafana/data';

import { LINKS_EXTENSION_POINT_ID } from '../../pages/AddedLinks';
import { testIds } from '../../testIds';

import { App } from './components/App';

export const plugin = new AppPlugin<{}>()
  .setRootPage(App)
  .configureExtensionLink({
    title: 'Open from B',
    description: 'Open a modal from plugin B',
    extensionPointId: 'plugins/grafana-extensionstest-app/actions',
    onClick: (_, { openModal }) => {
      openModal({
        title: 'Modal from app B',
        body: () => <div data-testid={testIds.appB.modal}>From plugin B</div>,
      });
    },
  })
  .configureExtensionComponent({
    extensionPointId: 'plugins/grafana-extensionexample2-app/configure-extension-component/v1',
    title: 'Configure extension component from B',
    description: 'A component that can be reused by other app plugins. Shared using configureExtensionComponent api',
    component: ({ name }: { name: string }) => <div data-testid={testIds.appB.reusableComponent}>Hello {name}!</div>,
  })
  .addComponent<{ name: string }>({
    targets: 'plugins/grafana-extensionexample2-app/addComponent/v1',
    title: 'Added component from B',
    description: 'A component that can be reused by other app plugins. Shared using addComponent api',
    component: ({ name }: { name: string }) => (
      <div data-testid={testIds.appB.reusableAddedComponent}>Hello {name}!</div>
    ),
  })
  .addLink({
    title: 'Open from B',
    description: 'Open a modal from plugin B',
    targets: [LINKS_EXTENSION_POINT_ID],
    onClick: (_, { openModal }) => {
      openModal({
        title: 'Modal from app B',
        body: () => <div data-testid={testIds.appB.modal}>From plugin B</div>,
      });
    },
  });
