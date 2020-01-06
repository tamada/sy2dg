'use strict'

const color = (d) => {
    const scale = d3.scaleLinear()
        .domain([1, 4])
        .range(['red', 'blue'])
    return scale(parseInt(d.group))
}

const radius = (d) => {
    if (d == undefined) {
        return 4
    }
    return 4
}

const drag = (simulation) => {
    const dragstarted = (d) => {
        if (!d3.event.active) {
            simulation.alphaTarget(0.3).restart()
        }
        d.fx = d.x
        d.fy = d.y
    }

    const dragged = (d) => {
        d.fx = d3.event.x
        d.fy = d3.event.y
    }

    const dragended = (d) => {
        if (!d3.event.active) {
            simulation.alphaTarget(0)
        }
        d.fx = null
        d.fy = null
    }

    return d3.drag()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended)
}

const labelLocationX = (d) => {
    return d.x + radius(d) + 2
}

const labelLocationY = (d) => {
    return d.y
}

const draw_graph_impl = (dataset) => {
    const w = 640
    const h = 640


    const svg = d3.select("#viewBox")
    const arrow = svg.append("defs")
    const marker = arrow.append("marker")
        .attr("id", "arrow")
        .attr("markerWidth", "4")
        .attr("markerHeight", "4")
        .attr("refX", "0")
        .attr("refY", "2")
        .attr("orient", "auto")
    marker.append("path")
        .attr("d", "M 0,0 V 4 L 4,2 Z")
        .attr("fill", "red")

    const node = svg.append("g")
        .selectAll("circle")
        .data(dataset.nodes)
        .join("circle")
        .attr("fill", color)

    d3.selectAll("circle")
        .attr('r', radius)

    const link = svg.append("g")
        .selectAll("line")
        .data(dataset.edges)
        .join("path")
        .attr("stroke", "#999")
        .attr("stroke-opacity", 0.6)
        .attr("stroke-width", 1)
        .attr("fill", "none")
        .attr("marker-end", "url(#arrow)")

    const labels = svg.append("g")
        .selectAll("text")
        .data(dataset.nodes)
        .join("text")
        .attr("class", "label")
        .attr("fill", "black")
        .attr("dominant-baseline", "middle")
        .text(d => d.name)

    // adjustment the directed line
    // refer from https://qiita.com/daxanya1/items/734e65a7ca58bbe2a98c
    const path = svg.selectAll("path")
    const totalLength = path.node().getTotalLength()
    svg.selectAll("path")
        .attr('stroke-dasharray', `0 4 ${totalLength} 4`)
        .attr("stroke-dashoffset", 0)

    const simulation = d3.forceSimulation(dataset.nodes)
        .force("link", d3.forceLink(dataset.edges).id(d => d.id))
        .force("charge", d3.forceManyBody())
        .force("center", d3.forceCenter(320, 320))
        .force("x", d3.forceX())
        .force("y", d3.forceY())
        .on("tick", () => {
            link
                .attr("d", (d) => `M${d.source.x},${d.source.y},${d.target.x},${d.target.y}`)
            node
                .attr("cx", (d) => d.x)
                .attr("cy", (d) => d.y)
            labels
                .attr("dx", labelLocationX)
                .attr("dy", labelLocationY)
        })
    node.call(drag(simulation))
}

const draw_graph = () => {
    d3.json("subdataset.json")
        .then((data) => draw_graph_impl(data))
        .catch((error) => d3.select("#message").text(error))
}
