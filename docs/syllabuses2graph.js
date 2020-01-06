'use strict'

const createCheckbox = (id, value, className) => {
    const input = document.createElement('input')
    input.type = 'checkbox'
    input.value = value
    input.id = id
    input.checked = true
    input.setAttribute('class', className)
    return input
}

const insertCheckbox = (item, uls) => {
    const label = document.createElement('label')
    label.appendChild(createCheckbox(`checkbox_${item.id}`, item.id, 'lecture-names'))
    label.appendChild(document.createTextNode(`${item.name}`))

    const li = document.createElement('li')
    li.appendChild(label)
    uls[parseInt(item.group) - 1].appendChild(li)
}

const createUl = (id) => {
    const ul = document.createElement('ul')
    ul.id = id
    return ul
}

const updateGradeMenu = (id, gradeCheck) => {
    const subULId = id.slice(0, 3)
    const subul = document.getElementById(subULId)
    const inputs = Array.from(subul.getElementsByTagName('input'))
    let nextFlag = false
    if (!gradeCheck.indeterminate && gradeCheck.checked) {
        nextFlag = true
    }
    inputs.forEach(input => input.checked = nextFlag)
    gradeCheck.checked = nextFlag
    gradeCheck.indeterminate = false

    network.body.data.nodes.update(inputs.map(input => { return { id: input.value, hidden: !input.checked } }))
}

const updateEvent = (network) => {
    const inputs = document.getElementsByClassName('lecture-names')
    Array.from(inputs).forEach(lectureInput => {
        lectureInput.addEventListener('click', () => toggleVisibleOfLecture(network, lectureInput))
    })

    const grades = document.getElementsByClassName('grade-menu')
    Array.from(grades).forEach(input => {
        input.addEventListener('click', () => updateGradeMenu(input.id, input))
    })
}

const findParentElementByTagName = (element, tagName) => {
    const name = tagName.toUpperCase()
    let e = element
    for (; e != null; e = e.parentElement) {
        if (e.tagName.toUpperCase() === name) {
            return e
        }
    }
    return null
}

const updateGradeCheck = (input) => {
    const parentUl = findParentElementByTagName(input.parentElement, 'ul')
    const id = parentUl.id
    const parentInput = document.getElementById(`${id}-grade`)
    const inputs = Array.from(parentUl.getElementsByTagName('input'))
    parentInput.checked = inputs.every(input => input.checked == true)
    if (!parentInput.checked) {
        parentInput.indeterminate = inputs.some(input => input.checked == true)
    }
}

const toggleVisibleOfLecture = (network, input) => {
    network.body.data.nodes.update({ id: input.value, hidden: !input.checked })
    updateGradeCheck(input)
}

const addSubUl = (ul, subul) => {
    const li = document.createElement('li')
    const label = document.createElement('label')
    const input = createCheckbox(`${subul.id}-grade`, `${subul.id}-grade-checks`, 'grade-menu')

    label.appendChild(document.createTextNode(`${subul.id} grade`))
    label.appendChild(input)

    const details = document.createElement('details')
    const summary = document.createElement('summary')
    summary.appendChild(label)
    details.appendChild(summary)
    details.appendChild(subul)
    // li.appendChild(label)
    // li.appendChild(subul)
    ul.appendChild(details)
}

const listup = (network) => {
    const uls = [createUl('1st'), createUl('2nd'), createUl('3rd'), createUl('4th')]

    dataset.nodes.forEach(item => insertCheckbox(item, uls))
    const div = document.getElementById("lectures")
    div.textContent = null
    const ul = document.createElement('ul')
    uls.forEach(subul => addSubUl(ul, subul))
    div.appendChild(ul)

    updateEvent(network)
    return div
}

const init_sy2dg = () => {
    const network = draw_graph()
    listup(network)
}
