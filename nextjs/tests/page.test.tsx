import { render, screen } from '@testing-library/react';
import { expect, test } from 'vitest';
import Home from '@/app/page';

test('the home page renders', () => {
    render(<Home />);

    expect(screen.getByRole('heading', { name: 'Hello from Laravel Cloud' })).toBeDefined();
});
