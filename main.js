import { html, ReactiveElement } from 'cami';
import { json } from 'd3-fetch';
import * as Plot from '@observablehq/plot';

class BarChart extends ReactiveElement {
  
  title = '';
  subtitle = '';
  caption = '';
  src = '';
  plotData = [];

  fetchData() {
    this.plotData = this.query({
      queryKey: [ 'plotData' ],
      queryFn: () => json(this.src),
      staleTime: Infinity
    });
  }

  onConnect() {
    this.observableAttributes({
      title: (title) => title,
      subtitle: (subtitle) => subtitle,
      caption: (caption) => caption,
      src: (src) => src
    });

    this.fetchData();
  }

  template() {
    var plt;
    if (this.plotData.data) {
      plt = Plot.plot({
        title: this.title,
        subtitle: this.subtitle,
        caption: this.caption,
        marks: [
          Plot.barY(this.plotData.data, { x: "Name", y: "Count", sort: { x: "y"} }),
          Plot.ruleY([0])
        ]
      })
    }
    return html`
    ${plt}
   `;
  }
}

customElements.define('bar-chart', BarChart);