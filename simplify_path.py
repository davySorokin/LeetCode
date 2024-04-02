class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = []
        components = path.split('/')

        for component in components:
            if component == "..":
                if stack:
                    stack.pop()
            elif component and component != ".":
                stack.append(component)
        
        # Join the components in the stack with '/' to form the simplified path
        return "/" + "/".join(stack)
