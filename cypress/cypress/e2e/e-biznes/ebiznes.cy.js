describe('products', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173/')
  })

  it('loads', () => {})

  it('retrieves products', () => {
    cy.intercept('GET', '/products', {
      statusCode: 200,
      body: [
        { id: 1, name: 'Product 1', price: 100 },
        { id: 2, name: 'Product 2', price: 200 }
      ]
    }).as('getProducts')

    cy.visit('http://localhost:5173/')
    cy.wait('@getProducts').its('response.statusCode').should('eq', 200)
  })

  it('contains title', () => {
    cy.get('h1').should('contain', 'Products')
  })

  it('contains products', () => {
    cy.get('.product-card').should('have.length.greaterThan', 0)
  })

  it('selects product', () => {
    cy.get('.product-card').first().click()
    cy.get('input[type="checkbox"]').first().should('be.checked')
  })

  it('contains title', () => {
    cy.get('.product-name').first().should('exist')
  })

  it('contains price', () => {
    cy.get('.product-price').first().should('exist')
  })

  it('contains checkout button', () => {
    cy.get('.checkout-button').should('exist')
  })

  it('checkout button forwards to checkout when products selected', () => {
    cy.get('.product-card').first().click()
    cy.get('.checkout-button').first().click()
    cy.url().should('include', '/checkout')
  })

  it('checkout button doesn\'t forward to checkout when no products selected', () => {
    cy.url().should('not.include', '/checkout')
  })
})

describe('checkout', () => {
  beforeEach(() => {
    cy.visit('http://localhost:5173/')
    cy.get('.product-card').first().click()
    cy.get('.checkout-button').first().click()
  })

  it('contains title', () => {
    cy.get('h1').should('contain', 'Checkout')
  })

  it('contains product list', () => {
    cy.get('.checkout-list').should('exist').should('have.length.greaterThan', 0)
  })

  it('list contains products', () => {
    cy.get('.checkout-list').should('have.length.greaterThan', 0)
    cy.get('.checkout-item').should('exist')
  })

  it('contains total price', () => {
    cy.get('.total-price').should('exist')
  })

  it('total price is correct', () => {
    cy.get('.total-price').invoke('text').then((text) => {
      const price = parseFloat(text.replace(/[^0-9.-]+/g, ''))
      expect(price).to.be.greaterThan(0)
    })
  })

  it('contains pay button', () => {
    cy.get('.submit-button').should('exist')
  })

  it('pay button has alert', () => {
    cy.get('.submit-button').click()
    cy.on('window:alert', (text) => {
      expect(text).to.contains('Payment successful!')
    })
  })

  it('successful payment redirects to home page', () => {
    cy.get('.submit-button').click()
    cy.url().should('include', '/')
  })

  it('displays no products selected when manually navigating to checkout', () => {
    cy.visit('http://localhost:5173/')
    cy.visit('http://localhost:5173/checkout')
    cy.get('.checkout-list').should('not.exist')
    cy.get('.checkout-item').should('not.exist')
    cy.get('.total-price').should('not.exist')
    cy.get('p').should('exist').contains('No products selected.')
  })

  it('sends payment request', () => {
    cy.intercept('POST', '/payments', {
      statusCode: 200,
      body: { success: true }
    }).as('paymentRequest')

    cy.get('.submit-button').click()
    cy.wait('@paymentRequest').its('response.statusCode').should('eq', 200)
  })
})