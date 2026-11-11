import XCTest
import SwiftTreeSitter
import TreeSitterDockerfileOrchard

final class TreeSitterDockerfileOrchardTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_dockerfile_orchard())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Dockerfile grammar")
    }
}
