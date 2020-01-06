'use strict'

const convertToVisNode = (item) => {
    const node = {
        id: item.id, label: item.name, group: item.group,
        value: 20 + item.weight * 3, url: item.url,
        semester: item.semester, teachers: item.teacherNames
    }
    return node
}

const convertToVisEdge = (item) => {
    return { from: item.source, to: item.target, value: item.weight, arrows: item.arrows }
}

const createOptions = () => {
    return {
        physics: true,
        interaction: {
            tooltipDelay: 200,
        },
        nodes: {
            shape: "dot",
            size: 20,
            borderWidth: 2
        },
        edges: {
            arrows: 'to',
            color: { inherit: 'from' },
            smooth: true
        },
        groups: {
            1: {
                color: {
                    background: "#97C2FC",
                    border: "#2B7CE9",
                    highlight: { background: "#D2E5FF", border: "#2B7CE9" },
                    hover: { background: "#D2E5FF", border: "#2B7CE9" }
                }
            },
            2: {
                color: {
                    background: "#FFFF00",
                    border: "#FFA500",
                    highlight: { background: "#FFFFA3", border: "#FFA500" },
                    hover: { background: "#FFFFA3", border: "#FFA500" }
                }
            },
            3: {
                color: {
                    background: "#FB7E81", border: "#FA0A10",
                    highlight: { background: "#FFAFB1", border: "#FA0A10" },
                    hover: { background: "#FFAFB1", border: "#FA0A10" }
                }
            },
            4: {
                color: {
                    background: "#7BE141",
                    border: "#41A906",
                    highlight: { background: "#A1EC76", border: "#41A906" },
                    hover: { background: "#A1EC76", border: "#41A906" }
                }
            }
        }
    }
}

const clearSelection = () => {
}

function neighbourhoodHighlight(params) {
    const node = nodesDataset.get(params.nodes[0])
    if (node != undefined) {
        console.log(`open ${JSON.stringify(node)}`)
    }
}

const buildNodeHTML = (node) => {
    const anchor = document.createElement("a")
    anchor.href = node.url
    anchor.target = "newweb"
    anchor.textContent = node.label
    return anchor
}

const createDlItem = (tagName, array) => {
    const item = document.createElement(tagName)
    if(Array.isArray(array)) {
        const ul = document.createElement('ul')
        array.forEach(entry => {
            const li = document.createElement('li')
            li.appendChild(entry)
            ul.appendChild(li)
        })
        item.appendChild(ul)
    }
    else {
        item.appendChild(array)
    }
    return item
}

const buildHTML = (node, befores, afters) => {
    const dl = document.createElement('dl')
    dl.appendChild(createDlItem('dt', document.createTextNode('Syllabus')))
    const dd = document.createElement('dd')
    dd.appendChild(buildNodeHTML(node))
    dd.appendChild(document.createTextNode(` (${node.group}年次生)`))
    dl.appendChild(dd)
    dl.appendChild(createDlItem('dt', document.createTextNode('Teachers')))
    dl.appendChild(createDlItem('dd', document.createTextNode(node.teachers.join("，"))))
    dl.appendChild(createDlItem('dt', document.createTextNode('Before')))
    dl.appendChild(createDlItem('dd', befores.map(item => buildNodeHTML(item))))
    dl.appendChild(createDlItem('dt', document.createTextNode('After')))
    dl.appendChild(createDlItem('dd', afters.map(item => buildNodeHTML(item))))
/*
    return `<dl>
    <dt>Syllabus</dt><dd>${buildNodeHTML(node)}（${node.group}年次生）</dd>
    <dt>Teachers</dt><dd>${node.teachers.join("，")}</dd>
    <dt>Before</dt><dd>${befores.map(item => buildNodeHTML(item)).join("，")}</dd>
    <dt>After</dt><dd>${afters.map(item => buildNodeHTML(item)).join("，")}</dd>
</dl>`
*/
    return dl
}

const findNodes = (nodeIds) => {
    return nodeIds.map(id => nodesDataset.get(id))
}

const findBeforeLectures = (node) => {
    return findNodes(dataset.edges.filter(item => item.target == node.id).map(edge => edge.source))
}

const findAfterLectures = (node) => {
    return findNodes(dataset.edges.filter(item => item.source == node.id).map(edge => edge.target))
}

const showLectureInfo = (params) => {
    const node = nodesDataset.get(params.nodes[0])
    if (node != undefined) {
        const div = document.getElementById('message')
        while(div.firstChild) div.removeChild(div.firstChild)
        div.appendChild(buildHTML(node, findBeforeLectures(node), findAfterLectures(node)))
    }
}

let network
const nodesDataset = new vis.DataSet(dataset.nodes.map(convertToVisNode))
const edgesDataset = new vis.DataSet(dataset.edges.map(convertToVisEdge))

const draw_graph = () => {
    const container = document.getElementById('graph')
    network = new vis.Network(container, { nodes: nodesDataset, edges: edgesDataset }, createOptions())

    network.on('click', (params) => {
        if (params.nodes.length == 0) {
            clearSelection()
        } else if (params.nodes.length == 1) {
            showLectureInfo(params)
            neighbourhoodHighlight(params)
        }
    })
    return network
}
